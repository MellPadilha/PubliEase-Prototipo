package handlers

import (
	"CrudProject/internal/configs"
	"CrudProject/pkg/models"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type HanResponse struct {
	Permissions []string `json:"permissions"`
}

type EuRoboResponse struct {
	Message string `json:"message"`
}

func VerifyApplication(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup   // WaitGroup para sincronizar as requisições
	done := make(chan bool) // channel para sinalizar conclusão das reqs

	// Variáveis para armazenar as responses
	var respHannibal HanResponse
	var respEuRobo EuRoboResponse
	var errHannibal, errEuRobo error
	var mutex sync.Mutex

	// setting form
	var form models.Form
	form.Name = r.FormValue("name")
	form.Package = r.FormValue("package")
	form.HasLocationAccess, _ = strconv.ParseBool(r.FormValue("hasLocationAccess"))
	form.LocationJustification = r.FormValue("locationJustification")
	form.HasCameraAccess, _ = strconv.ParseBool(r.FormValue("hasCameraAccess"))
	form.CameraJustification = r.FormValue("cameraJustification")

	// setting file
	file, header, err := r.FormFile("file")
	if err != nil {
		returnGenericError(w, "Form - Error: file not received or corrupted")
		return
	}
	defer file.Close()
	form.Apk = file

	// creating dir
	if err = os.MkdirAll("./uploads", os.ModePerm); err != nil {
		returnGenericError(w, "File - Error: problem creating folder")
		return
	}

	// saving file
	fileCreate, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		returnGenericError(w, "File - Error: no permission to write file")
		return
	}
	defer fileCreate.Close()
	io.Copy(fileCreate, file)

	// Open the saved file for sending
	form.Apk, err = os.Open("./uploads/" + header.Filename)
	if err != nil {
		returnGenericError(w, "File - Error: no permission to open file")
		return
	}
	defer form.Apk.Close()

	// setting request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	_ = writer.WriteField("name", form.Name)
	_ = writer.WriteField("package", form.Package)

	// Add the file
	fw, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		returnGenericError(w, "File - Error: problem preparing file for sending")
		return
	}

	_, err = io.Copy(fw, form.Apk)
	if err != nil {
		returnGenericError(w, "File - Error: problem making file available for sending")
		return
	}
	writer.Close()

	// Adicionando goroutines para processamento paralelo
	wg.Add(2)

	// Requisição para Hannibal
	go func() {
		defer wg.Done()
		req, err := http.NewRequest("POST",
			configs.GetHannibalHost()+configs.GetHannibalRoute(),
			&requestBody)
		if err != nil {
			mutex.Lock()
			errHannibal = err
			mutex.Unlock()
			done <- false
			return
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			mutex.Lock()
			errHannibal = err
			mutex.Unlock()
			done <- false
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			mutex.Lock()
			errHannibal = err
			mutex.Unlock()
			done <- false
			return
		}

		mutex.Lock()
		if err = json.Unmarshal(respBody, &respHannibal); err != nil { // Parse []byte to go struct pointer
			returnGenericError(w, "Emulator - Error: problem emulator response")
		}
		mutex.Unlock()

		done <- true
	}()

	// Requisição para EuRobo ou GAROTO
	go func() {
		defer wg.Done()
		motivo := models.Dados{
			Frase: form.LocationJustification,
		}

		bodyMotivo, err := json.Marshal(motivo)
		if err != nil {
			mutex.Lock()
			errEuRobo = err
			mutex.Unlock()
			done <- false
			return
		}

		reqeurobo, err := http.NewRequest("POST", configs.GetEuRoboHost()+configs.GetEuRoboRoute(), bytes.NewBuffer(bodyMotivo))
		if err != nil {
			mutex.Lock()
			errEuRobo = err
			mutex.Unlock()
			done <- false
			return
		}
		reqeurobo.Header.Set("Content-Type", "application/json")

		clientEurobo := &http.Client{}
		respEurobo, err := clientEurobo.Do(reqeurobo)
		if err != nil {
			mutex.Lock()
			errEuRobo = err
			mutex.Unlock()
			done <- false
			return
		}
		defer respEurobo.Body.Close()

		respBody, err := io.ReadAll(respEurobo.Body)
		if err != nil {
			mutex.Lock()
			errEuRobo = err
			mutex.Unlock()
			done <- false
			return
		}

		mutex.Lock()
		if err = json.Unmarshal(respBody, &respEuRobo); err != nil { // Parse []byte to go struct pointer
			returnGenericError(w, "Emulator - Error: problem emulator response")
		}
		mutex.Unlock()

		done <- true
	}()

	// Espera até que todos respondam
	go func() {
		wg.Wait()
		close(done)
	}()

	for range done {
		// O canal é fechado quando todas as respostas chegarem parando o for
	}

	// Depois que todas as responses forem concluídas, envia a resposta pro front
	mutex.Lock()
	defer mutex.Unlock()

	if errHannibal != nil {
		returnGenericError(w, "Emulator - Error: "+errHannibal.Error())
		return
	}
	if errEuRobo != nil {
		returnGenericError(w, "EuRobo - Error: "+errEuRobo.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"permissions": respHannibal.Permissions,
		"prediction":  respEuRobo.Message,
	})
}

func returnGenericError(w http.ResponseWriter, error string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]any{
		"error": error,
	})
}

import { useForm } from "@mantine/form";
import { useState } from "react";
import {
  Button,
  Group,
  TextInput,
  Grid,
  Switch,
  Text,
  Progress,
  Modal
} from "@mantine/core";
import { sendPostRequest } from "../../script/ApiRequest";
import { Dropzone } from "@mantine/dropzone";
import ReportModal from "../modalreport/report";

function DataForm() {
  const [sucessPred, setSucessPred] = useState(0);
  const [sucessLoc, setSucessLoc] = useState(0);
  const [opened, setModalOpened] = useState(false);
  const [progress, setProgress] = useState(0);
  const openRef = null;
  const [info, setInfo] = useState("");
  const [info2, setInfo2] = useState("");

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      package: "",
      hasLocationAccess: false,
      locationJustification: "",
      hasCameraAccess: false,
      cameraJustification: "",
      file: null,
    },
  });

  const handleSubmit = async (e) => {
    if (
      form.getValues().name &&
      form.getValues().package &&  
      form.getValues().file
    ) {

      try {
        setProgress(10)
        await espera(4000)
        setProgress(30)
        await espera(3000)
        setProgress(50)
        await espera(5000)
        setProgress(70)
        await espera(3000)
        let res = await sendPostRequest(form.getValues());
        console.log(res)
        setProgress(100);
        calcAprove(res.data)
      } catch (error) {
        console.error("Erro ao enviar o formulário:", error);
      } 
    }
  };

  function espera(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  function calcAprove(res){
      var location = res.permissions.includes("android.permission.ACCESS_FINE_LOCATION") || res.permissions.includes("android.permission.ACCESS_COARSE_LOCATION");

      if ( location && form.getValues().hasLocationAccess == true){
        setInfo("✅ As informações passadas batem com o informado pelo emulador")
        setSucessLoc(49);
      }
      else {
        setSucessLoc(0);
        setInfo("⚠️  Não foi informado corretamente os acessos as funcionalidades.")
      }

      if (res.prediction == "True") {
        setInfo2("✅ Seu motivo é um motivo valido de acordo com as diretrizes")
        setSucessPred(49)
      } 
      else {
        setSucessPred(0);
        setInfo2("⚠️ Não foi informado um motivo valido, de acordo com as diretrizes.");
      }

      setModalOpened(true);
      setProgress(0); // Abre o modal
  }
  return (
    <div>
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <Grid>
          <Grid.Col span={6}>
            <TextInput
              label="Nome do Aplicativo"
              placeholder="Digite o nome"
              withAsterisk
              key={form.key("name")}
              {...form.getInputProps("name")}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <TextInput
              label="Nome do Pacote (Package)"
              placeholder="Digite o nome do pacote"
              withAsterisk
              key={form.key("package")}
              {...form.getInputProps("package")}
            />
          </Grid.Col>
          <Grid.Col span={6}>
            <h4>Você solicita acesso a localização do dispositivo?</h4>
            <Switch
              size="xl"
              onLabel="Sim"
              offLabel="Não"
              style={{ margin: "20px" }}
              checked={form.getValues().hasLocationAccess}
              onClick={() =>
                form.setValues({
                  hasLocationAccess: !form.getValues().hasLocationAccess,
                })
              }
            />
            {form.getValues().hasLocationAccess && (
              <TextInput
                label="Motivo de acesso a localização"
                placeholder="Digite o motivo"
                key={form.key("locationJustification")}
                {...form.getInputProps("locationJustification")}
              />
            )}
          </Grid.Col>
          <Grid.Col span={6}>
            <Dropzone
              openRef={openRef}
              onDrop={(files) => {
                form.setFieldValue("file", files ? files[0] : null);
              }}
            >
              <Group
                justify="center"
                gap="xl"
                mih={220}
                style={{ pointerEvents: "none" }}
              >
                {!form.getValues().file && (
                  <div style={{ textAlign: "center" }}>
                    <Text size="xl" inline>
                      Clique para selecionar o arquivo APK
                    </Text>
                    <Text size="sm" c="dimmed" inline mt={7}>
                      Apenas um arquivo pode ser selecionado.
                    </Text>
                    <Text size="sm" c="dimmed" inline mt={7}>
                      Apenas arquivos "apk" são permitidos.
                    </Text>
                  </div>
                )}
                {form.getValues().file && (
                  <div style={{ textAlign: "center" }}>
                    <Text size="xl" inline>
                      {form.getValues().file.name}
                    </Text>
                    <Text size="sm" c="dimmed" inline mt={7}>
                      Para alterar, clique e selecione um novo aplicativo.
                    </Text>
                  </div>
                )}
              </Group>
            </Dropzone>
          </Grid.Col>

          <Grid.Col span={12}>
            <Group justify="center" mt="md">
              <Button type="submit">
                Enviar
              </Button>
            </Group>
             <div style={{ marginTop: "20px" }}>
                <Progress color="blue" size="lg" value={progress} striped animated />
            </div>
          </Grid.Col>
        </Grid>
      </form>

      <Modal
        opened={opened}
        onClose={() => (setModalOpened(false), setSucessPred(0), setSucessLoc(0)) }
        title="Relatório do APK"
      >
        <p>
          Após analise for identificado que você tem {sucessLoc + sucessPred}% de chance de ser aprovado na avaliação da PlayStore.
          <br/>
          <br/>
            • {info}
          <br/>
          <br/>
            • {info2}
        </p>
      </Modal>
    </div>
  );
}

export default DataForm;

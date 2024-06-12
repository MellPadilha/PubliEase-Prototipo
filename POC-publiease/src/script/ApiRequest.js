import axios from "axios";
axios.defaults.headers.post['Content-Type'] ='application/x-www-form-urlencoded';

export async function sendPostRequest(form) {
    let response;
    await axios.post("http://localhost:9010/api/",
            form,
            {
                headers: {
                    "Content-type": "multipart/form-data, */*",
                    "Accept": "*/*",
                },
             },
        )
        .then(res => response = res)
        .catch(err => console.error(err))
    return response
  }



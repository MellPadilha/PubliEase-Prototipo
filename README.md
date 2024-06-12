# PubliEase-Prototipo
Unificação de todos os projetos que compõem o prototipo do publiease

# Procedimento de execução dos projetos

Este documento descreve os passos necessários para configurar o ambiente e validar um aplicativo Android utilizando diversas ferramentas e tecnologias, incluindo Docker, Go, Python, e npm.

## Requisistos

- **Docker**: Deve estar previamente configurado e conectado ao servidor Docker Hub.
- **Go**: Deve estar incluído nas variáveis de ambiente para que seu acesso seja permitido por todo o ambiente.
- **Python**: Deve estar incluído nas variáveis de ambiente para que seu acesso seja permitido por todo o ambiente.
- **Biblioteca "pip"**: Pertencente ao Python, deve estar instalada.
- **Gerenciador de pacotes "npm"**: Deve estar instalado e disponível no ambiente.

## Informações necessárias

- **Aplicativo Android (APK)**: Deve estar disponível para ser testado no emulador. O APK deve ser compatível com o Android 7.0 ou superior.
- **Nome e Package do Aplicativo**: O Package é fundamental para que o emulador consiga interagir com o software enviado para validação.

## Procedimento para Validação

1. **Realizar Download do Código Fonte**:
   - Faça o download do código fonte do projeto.
   - Cada projeto tem seu readme interno caso seja necessário.

2. **Descompactar o Projeto**:
   - Extraia os arquivos do projeto para o diretório de sua escolha.

3. **Inicializar o Serviço Docker**:
   - Inicie o serviço Docker para que as demais aplicações possam acessá-lo.

4. **Abrir o Terminal**:
   - Execute todos os passos seguintes no terminal.

5. **Instalação dos Pacotes do Projeto "POC-publiease"**:
   - Acesse o diretório do projeto "POC-publiease":
     ```bash
     cd POC-publiease
     ```
   - Execute o comando para instalar os pacotes:
     ```bash
     npm install
     ```

6. **Iniciar o Projeto "kratos-api"**:
   - Acesse o diretório do projeto "kratos-api":
     ```bash
     cd ../api_kratos
     ```
   - Execute o comando para baixar as dependências e iniciar o projeto:
     ```bash
     go run cmd/*.go
     ```

7. **Preparar o Ambiente do Emulador Android no Projeto "hannibal"**:
   - Acesse o diretório do projeto "hannibal":
     ```bash
     cd ../hannibal_api
     ```
   - Execute o comando para preparar o ambiente do emulador:
     ```bash
     docker pull thyrlian/android-sdk
     ```

8. **Configurar Permissões e Caminho do SDK**:
   - Execute o comando para obter o nome do usuário:
     ```bash
     whoami
     ```
   - Anote o nome do usuário e, em seguida, execute o comando para alterar as permissões do SDK:
     ```bash
     chown <nome_do_usuario> sdk/
     ```
   - Estando na pasta do emulador, execute o comando para obter o caminho absoluto:
     ```bash
     pwd
     ```
   - Anote o resultado e edite o arquivo `config.yaml` na linha 20, substituindo pelo caminho anotado e adicionando o texto "/sdk:/opt/android-sdk". O resultado deve ser parecido com:
     ```
     /home/user/hannibal_api/sdk:/opt/android-sdk
     ```

9. **Iniciar o Projeto "hannibal"**:
   - Execute o comando para baixar as dependências e iniciar o projeto:
     ```bash
     go run cmd/hannibal/*.go
     ```

10. **Instalação de Dependências da IA**:
    - Acesse o diretório do projeto da IA:
      ```bash
      cd ../ApiEuRobo
      ```
    - Execute o comando para instalar as dependências:
      ```bash
      
      PARA WINDOWS: 
          python -m venv sklearn-env
          sklearn-env\Scripts\activate
          pip install -r requirements.txt
      PARA LINUX:
          python3 -m venv sklearn-env
          source sklearn-env/bin/activate 
          pip3 install -r requirements.txt
      ```

11. **Iniciar o Servidor do EuRobo (IA)**:
    - Execute o comando para iniciar o servidor:
      ```bash
      python3 main.py
      ```

12. **Iniciar o Servidor do Projeto "POC-publiease"**:
    - Volte ao diretório do projeto "POC-publiease":
      ```bash
      cd ../POC-publiease
      ```
    - Execute o comando para iniciar o servidor:
      ```bash
      npm run start
      ```

13. **Acessar o Servidor Local**:
    - Abra um navegador e acesse:
      ```
      localhost:3000
      ```

14. **Enviar o APK para Validação**:
    - Preencha o nome e package do aplicativo.
    - Clique no botão para solicitar acesso à localização do dispositivo.
    - Selecione o APK para enviar ao servidor.
    - Aguarde a conclusão do processo.
    - EXEMPLO DE DADOS PARA TESTE:
        - Nome: Twitter ou X
        - Package: com.twitter.android
        - Apk:Disponibilizamos alguns apks para download [aqui](https://drive.google.com/drive/folders/1cmjrM92v0BlIt-jkCEy_wr26WDefL9KW?usp=sharing) (O twitter se chama X.apk)
        - Motivo valido: Recomendar atividades de lazer ao usuários ou personalizar recomendação de conteúdo com base na localização
        - Motivo invalido: Roubar dados do usuarios ou Porque o aplicativo é meu

15. **Verificação dos Resultados**:
    - Verifique os dados exibidos na tela para garantir que o processo foi concluído corretamente.

## Conclusão

Siga todos os passos cuidadosamente para garantir que a configuração e validação do aplicativo sejam realizadas com sucesso. Se ocorrer algum problema, verifique os logs e os detalhes das configurações para identificar e corrigir possíveis erros.

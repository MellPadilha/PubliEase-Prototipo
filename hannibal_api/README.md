## Para executar este projeto

É necessário ter o go instalado e configurado para rodar globalmente na maquina e o docker estar instalado.
´Recomendamos rodar o projeto no Linux ou MAC´

1. Executar: docker pull thyrlian/android-sdk
2. Executar docker run -it --rm -v $(pwd)/sdk:/sdk thyrlian/android-sdk bash -c 'cp -a $ANDROID_HOME/. /sdk'
3. Executar o comando pwd dentro da pasta do projeto e copiar resultado
3. Alterar config.yaml: "/home/usr/go/src/hannibal/sdk::/opt/android-sdk" -> 
    "<resultado do pwd>/hannibal/sdk::/opt/android-sdk"
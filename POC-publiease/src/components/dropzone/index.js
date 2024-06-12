import { Button, Group, Text } from '@mantine/core';
import { Dropzone } from '@mantine/dropzone';

function DropzoneApk() {
  const openRef = null;

  return (
    <>
      <Dropzone openRef={openRef} onDrop={() => {}} accept={'application/vnd.android.package-archive'}>
          <Group justify="center" gap="xl" mih={220} style={{ pointerEvents: 'none' }}>
            <div>
              <Text size="xl" inline>
                Clique para selecionar o arquivo APK
              </Text>
              <Text size="sm" c="dimmed" inline mt={7}>
                Selecione apenas um arquivo, o arquivo precisar ser um apk.
              </Text>
            </div>
          </Group>
      </Dropzone>
      
    </>
  );
}

export default DropzoneApk;
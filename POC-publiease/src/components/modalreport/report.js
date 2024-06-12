import React, { useState } from 'react';
import { Modal, Button } from '@mantine/core';

function ReportModal({opened, setModalOpened}) {

  return (
    <div>
      <Button  onClick={() => setModalOpened(true)}>Abrir Relatório do APK</Button>

      {/* style={{display: "None"}} */}
      
      <Modal
        opened={opened}
        onClose={() => setModalOpened(false) }
        title="Relatório do APK"
      >
        <p>
          Aqui está o conteúdo do relatório do APK. Você pode adicionar mais detalhes
          sobre o relatório aqui, como informações sobre o APK, estatísticas, etc.
        </p>
      </Modal>
    </div>
  );
}

export default ReportModal;

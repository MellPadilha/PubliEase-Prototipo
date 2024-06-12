import DataForm from './components/form';
import DropzoneApk from './components/dropzone';

function App() {
  const screenSize = {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: '100vh', /* altura total da viewport */
  };

  const contentSize = {
    maxWidth: '800px',
    width: '100%',
    padding: '20px',
  };
  return (
    <div style={screenSize} >
      <div style={contentSize}> 
        <h1>POC - PubliEase</h1>
         <DataForm/>
      </div>
    </div>

  );
}

export default App;

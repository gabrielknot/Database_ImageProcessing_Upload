import './App.css';
import ApiContextProvider from './components/apiContextProvider';
import ImageCropper from './components/imageCropper'

function App() {
  return (
    <div className="App">
      <ApiContextProvider>
        <ImageCropper/>
      </ApiContextProvider>
    </div>
  );
}

export default App;

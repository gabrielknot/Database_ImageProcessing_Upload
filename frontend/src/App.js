import './App.css';
import ImageCropper from './components/imageCropper'
import logo from './logo192.png'

function App() {
  return (
    <div className="App">
      <ImageCropper src = {logo}/>
    </div>
  );
}

export default App;

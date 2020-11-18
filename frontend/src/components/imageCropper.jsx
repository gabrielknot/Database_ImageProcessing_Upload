import React, {Component} from 'react';
import Cropper from 'react-cropper';
import 'cropperjs/dist/cropper.css';
import './styles/imageCropper.css'

export default class imageCropper extends Component {
    constructor(){
        super();
        this.state = {
            imgDist:''
        }
        this.imgRef = React.createRef()
    }
    _crop() {
        this.setState({ ...this.state,imgDist: this.cropper.getCroppedCanvas().toDataURL("image/png")});
    }

    onCropperInit(cropper) {
        this.cropper = cropper;
    }

    onDragEnd(){
        const cropperImageData = this.cropper.getCroppedCanvas().toDataURL("image/png")
        console.log(`the current data is altered from ${this.state.imgDist.length} to ${cropperImageData.length}`)
        this.setState({ ...this.state,imgDist: cropperImageData})    
    }

    componentDidMount(){
        const img = this.props.src

    }

    render() {
        return (
            <div className="container">
                <div className="image-cropper">
                        
                    <Cropper
                        src={this.props.src}
                        ref={this.imgRef}
                        style={{height: 180, width: 480}}
                        // Cropper.js options 
                        initialAspectRatio={16 / 9}
                        guides={false}
                        crop={this._crop.bind(this)}
                        onInitialized={this.onCropperInit.bind(this)}
                    />
                    <div role = "buttom" 
                         className="submit-buttom" 
                         onClick={_=>console.log('aaaaaaaaa')}>
                             submit
                         </div>
                </div>
                
            </div>
        );
    }
}

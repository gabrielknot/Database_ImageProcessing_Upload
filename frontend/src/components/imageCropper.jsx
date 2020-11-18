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
        const canvasImgDataURL = this.cropper.getCroppedCanvas().toDataURL("image/png")
        console.log(canvasImgDataURL)
        this.setState({ ...this.state,imgDist: canvasImgDataURL});
    }

    onCropperInit(cropper) {
        this.cropper = cropper;
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
                        onInitialized={this.onCropperInit.bind(this)}
                    />
                    <div role = "buttom" 
                         className="submit-buttom" 
                         onClick={this._crop.bind(this)}>
                             <span>submit </span>
                         </div>
                </div>
                
            </div>
        );
    }
}

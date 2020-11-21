import React, {Component} from 'react';
import Cropper from 'react-cropper';
import 'cropperjs/dist/cropper.css';
import './styles/imageCropper.css';

export default class imageCropper extends Component {
    constructor(props){
        super(props);
        
        this.imgRef = React.createRef()

    }

    _crop() {
        const canvasImgDataURL = this.cropper.getCroppedCanvas().toDataURL("image/png")
        this.props.setCurrentImage(canvasImgDataURL);
        this.props.handlePutCurrentImage(
            this.props.currentDatabaseName,
            this.props.currentImageIndex
        )
    }

    onCropperInit(cropper) {
        this.cropper = cropper;
    }
    
    render() {
        const open = this.context.currentImage !== '' && this.context.currentImage !== undefined 
        console.log(this.context.currentImage)
        return (
            
            <div className="container">
                {open &&
                <div className="image-cropper">
                        
                    <Cropper
                        src={this.props.currentImage}
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
                }
            </div>
        );
    }
}

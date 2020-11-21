import React,{useState, useRef}  from 'react';
import './styles/dropContainer.css'

export default function ImageDownload(props){
    const imgRef = useRef()
    const currentImage = props.currentImage
    const currentDatabaseName = props.currentDatabaseName
    const currentDatabaseID = props.currentDatabaseID
    const currentImageIndex = props.currentImageIndex
    const extention = (currentImage !=='')&&(currentImage !== undefined)?
                      currentImage.split(";")[0].split("/")[1]:''
    
    return (
        <div className="imageDownload"
             onClick={_=>imgRef.current.click()}
        >
             <a
                ref={imgRef}
                download={`database_${currentDatabaseName}_id${currentDatabaseID}_${currentImageIndex}.${extention}`}
                href={"data:"+currentImage}
             >
                 Download
             </a>
        </div>
    );
    }

    
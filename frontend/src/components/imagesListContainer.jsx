import React  from 'react';
import './styles/imagesList.css'

export default function ImageListContainer(props){
    const images = !(props.images?.isArray())||(props.images==undefined)? [props.images]:[]
    const imageIndex = props.imageIndex
    const currentDatabase = props.currentDatabase
    
    return (
        <div className="imageList">
            {images.map((element, key)=>{
                const active = key === imageIndex
                return(
                    <div className="celImageList">
                        <div
                            onClick={_=>props.setCurrentImageIndex(key)}
                            role="buttom"
                        >
                            <img className={"image"+ active? " active":" inactive"}
                                src={element}
                                width={active?"40":"32"} 
                                height={active?"22":"18"}/>
                        </div>
                        <div className="imageDeleteButtom" 
                             onClick={_=>props.handleRemoveImageInDataBase(currentDatabase,imageIndex)}
                             role="buttom">
                            DELETE
                        </div>
                    </div>
                )
            })}
             
        </div>
    );
    }

    
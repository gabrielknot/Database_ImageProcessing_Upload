import React, {Component, createContext} from 'react';
import axios from 'axios';
import ImageCropper from './imageCropper';
import DatabaseList from './databaseList';
import ImageDownload from './imageDownload';
import ImageUpload from './imageUpload';
import ImageListContainer from './imagesListContainer';
import DeleteDatabaseBuuttom from './deleteDatabaseButtom'


import './styles/apiContext.css'

const APIstate = {
    databases:[{id:-1,dbname:"None databases",images:[]}],
    currentImage:'',
    currentImageIndex:0,
    currentDatabase:{},
    currentDatabaseIndex:-1,
}


const Api = axios.create({
            baseURL:"http://localhost:3003/api/databases/"

        })
export default class apiContextProvider extends Component {
      
    constructor(){
        super();
        this.state = {...APIstate.data}
        this.refresh = this.refresh.bind(this) 

    }
    refresh(){
        Api.get("/")
        .then(resp => {
            if(resp.data != []){
                    this.setState({ 
                        ...this.state,
                        currentDatabase: this.state.currentDatabaseIndex>=0?
                        resp.data[this.state.currentDatabaseIndex]:{id:-1,dbname:"None databases",images:[]},
                        databases:resp.data
                })
            }
        })
        
    }
    handleAddImageInDataBase(database,img){
        const newDatabase = {...database,images: database.images.push(img)}
        Api.put(`/${newDatabase.id}`,newDatabase).then(_=>this.refresh())
    }
    handleRemoveImageInDataBase(database, imgKey){
        const newDatabase = {...database,images: database.images.splice(imgKey,1)}
        Api.put(`/${database.id}`,newDatabase).then(_=>this.refresh())
    }
    handlePutCurrentImage(currentDatabase,img,imgKey){
        const newDatabase = {...currentDatabase,images:currentDatabase.images.splice(imgKey,1,img)}
        Api.put(`/${newDatabase.id}`,newDatabase).then(_=>this.refresh())
    }
    handleAddDataBase(newDtabaseName){
        Api.post("/",{id:-1,dbname:newDtabaseName,images:[""]}).then(_=>this.refresh())
    }
    handleRemoveDataBase(database){
        Api.delete(`${database.id}`).then(_=>this.refresh())
    }
    setCurrentImage(imgDataURL){
        this.setState({
            ...this.state,
            currentImage:imgDataURL
        })
    }
    setCurrentImageIndex(imgKey){
        this.setState({
            ...this.state,
            currentImageIndex:imgKey
        })
    }
    setCurrentImageIndex(databaseKey){
        this.setState({
            ...this.state,
            currentDatabaseIndex:databaseKey
        })
    }
    componentDidMount(){
        console.log(this.state)
        Api.get("/")
        .then(resp => {
            if(resp.data != []){
                    this.setState({ 
                        ...this.state,
                        currentDatabase: this.state.currentDatabaseIndex>=0?
                        resp.data[this.state.currentDatabaseIndex]:{id:-1,dbname:"None databases",images:[]},
                        databases:resp.data
                })
            }
        })
    }
        
    render() {
        return (
            <div className='api-provider'>
                <div className="leftBar">                    
                    {{/* <DatabaseList
                        databases={_=>this.state.databases}
                        currentDatabaseID={this.state.currentDatabase?.ID}
                        handleAddDataBase={(newDtabaseName)=>{this.handleAddDataBase(newDtabaseName)}}
                        setCurrentDatabase={(databaseKey)=>{this.setCurrentImageIndex(databaseKey)}}
                    /> */}}
                </div>
                <div className="middlePage">
                    <div className="buttons">                        
                        <ImageUpload
                            currentDatabase={this.state.currentDatabase}
                            handleAddImageInDataBase={(database,img)=>{this.handleAddImageInDataBase(database,img)}}
                        />
                        <ImageDownload
                            currentImage={this.state.currentImage}
                            currentImageIndex={this.state.currentImageIndex}
                            currentDatabaseName={this.state.currentDatabase?.Dbname}
                            currentDatabaseID={this.state.currentDatabase?.id}
                        />
                    </div>
                    <ImageCropper 
                        currentImage={this.state.currentImage}
                        currentImageIndex={this.state.currentImageIndex}
                        currentDatabase={this.state.currentDatabase}
                        setCurrentImage={(imgDataURL)=>{this.setCurrentImage(imgDataURL)}}
                        handlePutCurrentImage={(currentDatabase,img,imgKey)=>{this.handlePutCurrentImage(currentDatabase,img,imgKey)}}
                        />
                    
                </div>
                <div className="rightBar">
                    <ImageListContainer
                        images={this.state.currentDatabase?.Images}
                        imageIndex={this.state.currentImageIndex}
                        currentDatabase={this.state.currentDatabase}
                        setCurrentImageIndex={(imgKey)=>{this.setCurrentImageIndex(imgKey)}}
                        handleRemoveImageInDataBase={(database,databaseKey)=>{this.handleRemoveImageInDataBase(database,databaseKey)}}
                    />
                    <DeleteDatabaseBuuttom
                        currentDatabaseID={this.state.currentDatabase?.id}
                        handleRemoveDataBase={(database)=>{this.handleRemoveDataBase(database)}}
                    />
                </div>
            </div> 
        );
    }
}

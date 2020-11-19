import React, {Component, createContext} from 'react';
import axios from 'axios'

const APIstate = {
    databases:[],
    currentImage:'',
    handleAddImageInDataBase:_=>{},
    handleRemoveImageInDataBase:_=>{},
    handleAddDataBase:_=>{},
    handleRemoveDataBase:_=>{},
    setCurrentImage:_=>{}

}

const contextAPI = createContext(APIstate)

const URLapi = "http://localhost:3000/api/databases"

export default class apiContextProvider extends Component {
    constructor(){
        super();
        this.state = APIstate

        this.refresh = this.refresh.bind(this) 

        
        this.handleAddImageInDataBase = this.this.handleAddImageInDataBase.bind(this)
        this.handleRemoveImageInDataBase = this.this.handleRemoveImageInDataBase.bind(this)
        
        this.handleAddDataBase = this.handleAddDataBase.bind(this)
        this.handleRemoveDataBase = this.handleRemoveDataBase.bind(this)

    }

    refresh(){
        axios.get(`${URLapi}`)
         .then(resp => this.refreshState(URLapi,resp))
        
    }
    
    refreshState(resp) {
        
        return this.setState({ ...this.state, databases: resp.data })
       
    }
    
    handleAddDataBase(){

        axios.post(URLapi, this.state.databases).then(_=> this.refresh())
    }

    handleRemoveDatabase(database){
    
        axios.delete(`${URLapi}/${database.id}`).then(_=> this.refresh()).catch(_=> this.refresh())
    }


    handleAddImageInDataBase(database,img){
        this.setState({...this.state,databases: this.state.databases
            .map(element => {
                if(element === database){
                    
                    element.images = element.images.push(img)
                }
                return element
            })
        });

        axios.put(`${URLapi}/${database.id}`).then(_=> this.refresh()).then(_=> this.refresh())
    }

    handleRemoveImageInDataBase(database,key){
        this.setState({...this.state,databases: this.state.databases
            .map(element => {
                if(element === database){
                    
                    element.images = element.images.splice(key,1)
                }
                return element
            })
            });
        
        axios.put(`${URLapi}/${database.id}`).then(_=> this.refresh()).catch(_=> this.refresh())
    }
    
    setCurrentImage(img){
        this.setState({...this.state, currentImage: img})
    }

    componentDidMount(){
        this.setState({
            ...this.state,
            handleAddImageInDataBase: this.handleAddImageInDataBase,
            handleRemoveImageInDataBase: this.handleRemoveImageInDataBase,
            handleAddDataBase: this.handleAddDataBase,
            handleRemoveDataBase: this.handleRemoveDataBase        
            })
        this.refresh()
    }    

    render() {
        return (
            <contextAPI.Provider value={this.state}>
                {this.props.children}
            </contextAPI.Provider>    
        );
    }
}

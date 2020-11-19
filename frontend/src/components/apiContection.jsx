import React, {Component, createContext} from 'react';
import axios from 'axios'

const APIstate = {
    databases:{
        testDatabase0:{
            id:1,
            dbname:'database0',
            images:[{
                imgTitle:'testImage2',
                imgBinBase64:"data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7                        ",
                date: new Date.now()
                }
            ]

        },
        testDatabase1:{
            id:2,
            dbname:'database1',
            images:[{
                imgTitle:'testImage2',
                imgBinBase64:"data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7                        ",
                date: new Date.now()
                }
            ]
        },
        testDatabase2:{
            id:3,
            dbname:'database2',
            images:[{
                imgTitle:'testImage2',
                imgBinBase64:"data:image/png;base64,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",
                date: new Date.now()
                }
            ]
        }
    }
}

const contextAPI = createContext(APIstate)

const URLapi = "http://localhost:3000/api/databases"
export default class apiContextProvider extends Component {
    constructor(){
        super();
        this.state = APIstate
        this.refresh = this.refresh.bind(this) 
        
        this.handleRemoveImg = this.handleRemoveImg.bind(this)
        this.handleRemoveImg = this.handleRemoveImg.bind(this)
        this.handleAddImage = this.handleAddImage.bind(this)
        this.handleAddDataBase = this.handleAddDataBase.bind(this)

    }
    refresh(){
        axios.get(`${URLapi}`)
         .then(resp => this.refreshState(URLapi,resp))
        
    }
    
    refreshState(URL,resp) {
        
        return this.setState({ ...this.state, databases: resp.data })
       
    }
    
    handleAddImage(database, name = "Undefined",imageBase64){
        
        axios.post(URLapi, {groups}).then(_=> this.refresh())
        
    }
     
    handleRemoveImg(database, name = "Undefined",imageBase64){
    
        axios.delete(`${URLapi}/${img.id}`).then(_=> this.refresh()).catch(_=> this.refresh())
    }

    handleAddDataBase(){
        
        axios.post(URLapi, {groups}).then(_=> this.refresh())
        
    }

    handleRemoveDatabase(img){
    
        axios.delete(`${URLapi}/${img.id}`).then(_=> this.refresh()).catch(_=> this.refresh())
    }

    componentDidMount(){
        this.refresh()
    }    

    render() {
        return (
            <div>
                
                
            </div>
        );
    }
}

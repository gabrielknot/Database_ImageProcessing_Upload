import React,{useState, useRef}  from 'react';



export default function DatabaseList(props){
    const databases = props.database
    const mapOfDatabaseNames = databases?.map(element=>{
        return element.dbname
    })

    const refInput = useRef()
    const [openInput, setOpenInput ] = useState(false)
    
    const [input, setInput] = useState('')

    const [invalidInput, setInvalidInput] = useState(false)
    
    const handleInput = (e)=>{
        const currentInput = e.target.value
        setInput(currentInput)
    
        if (mapOfDatabaseNames.indexOf(currentInput) >= 0){
            setInvalidInput(true)
        }
    }
    return (
        <div className="databaseList">
            {databases.map(database=>{
                const databaseName = database.dbname
                const databaseID = database.id
                const isCurrentDatabase = databaseID === props.currentDatabaseID

                return(
                    <div className={"databaseSelectionButtom"+isCurrentDatabase?"-current":""}
                         key={databaseID}
                         role="buttom"
                         onClick={_=>props.setCurrentDatabase(database)}
                         >
                             {databaseName}
                    </div>
                )
            })}
            {openInput &&
                <div>
                    <input 
                        ref={refInput}
                        onSubmit={(e)=>{
                            if (!invalidInput){
                                props.handleAddDataBase(input)
                                setOpenInput(false)
                            }
                        }}
                        onChange={e=> handleInput(e)}
                        placeholder="Database Name"
                    />{
                        invalidInput && 
                        <span className="invalidInput">Please Select a Valid Name</span>
                    }
                    <div className="submitButtom" 
                         role="buttom"
                         onClick={_=>{
                             if (!invalidInput){
                                 props.handleAddDataBase(input)
                             }
                         }}
                         >
                        submit
                    </div>
                </div>
            }{!openInput &&
                <div className="newDatabaseContainer">
                    <div className="newDatabaseButtom"
                        role="buttom"
                        onClick={_=>{
                            setOpenInput(!openInput)
                            refInput.current.focus()
                        }}
                    >
                        New Database
                    </div>
                </div>
            }
        </div>
    );
    }

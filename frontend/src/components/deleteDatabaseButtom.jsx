import React,{useState, useRef}  from 'react';
import './styles/deleteDatabaseButtom.css'

export default function DeleteDatabaseButtom(props){
    
    const currentDatabase = props.currentDatabase
    
    return (
        <div className="deleteDatabase">
            <div className="deleteDatabaseButtom"
                onClick={_=>{
                    const confirm = window.confirm(`You chose delete the ${currentDatabase.Dbname} database. Are you sure?`)
                    if(confirm){
                        props.handleRemoveDataBase(currentDatabase)
                     }
               }}
            >
                    DELETE DATABASE
            </div>
        </div>
    );
    }

    
import React from "react"
import { Outlet } from 'react-router-dom'
import axios from "axios"
import EditFlashcard from "../editFlashcard/editFlashcard"

import './editDeckRoute.css'

const EditDeckRoute = ({deckNames, deckID}) => {
    const url = "http://localhost:8080/api/deck/" + deckID
    
    const [deckInfo, setDeckInfo] = React.useState([])

    React.useEffect(() => {
        axios.get(url).then((response) => {
            setDeckInfo(response.data)
        })
    }, [url])

    function newCard() {
        const tmpDeck = new Map()
        tmpDeck.set("topic", "")
        tmpDeck.set("description", "")
        tmpDeck.set("id", deckInfo.length++)

        // create brand new array. loop through deckInfo and add elements to new array
        let tmpArr = []
        deckInfo.forEach((card) => {
            tmpArr.push(card)
        })
        tmpArr.push(tmpDeck)

        setDeckInfo(tmpArr)
    }

    function handleDeckInfoChange(index, e) {
        setDeckInfo(prev => {
            const newState = prev.map(card => {
                if (card.id === index) return {
                    ...card,
                    [e.target.name]: e.target.value
                }
                return card
            })
            return newState
        })
    }

    function saveCards() {
        axios.post(url, deckInfo, { params:{"deckName": deckNames.get(deckID)}})
    }
    
    return (
        <>
            <div className="wrapper">
                <div className="header">
                    <h1>Edit Flashcards</h1>
                    <button onClick={newCard} >Add Card</button>
                    <button onClick={saveCards} >Save Cards</button>
                </div>
                <ul className="flashcards">
                    {deckInfo.map((flaschard, num) => {
                                return (
                                    <li key={num} className="edit-flashcard-wrapper"> 
                                        <EditFlashcard flaschard={flaschard} editState={handleDeckInfoChange} />
                                    </li>
                                )
                            }
                        )}
                </ul>

            </div>
            <div id="detail">
                <Outlet />
            </div>
        </>
    )
}

export default EditDeckRoute;
import React from 'react'
import { Outlet } from 'react-router-dom'
import axios from 'axios'
import FlashcardParent from '../flashcard/flashcardParent'

const ReviewDeckRoute = ({deckID}) => {
    const url = (import.meta.env.VITE_API_URL || '') + '/api/deck/' + deckID;
    //const url = 'http://localhost:8080/api/deck/' + deckID

    const [deckInfo, setDeckInfo] = React.useState([])

    React.useEffect(() => {
        axios.get(url).then((response) => {
            console.log('reviewDeckRoute.jsx url response: ', response)
            setDeckInfo(response.data)
        })
    }, [url])

    console.log("reviewDeckRoute.jsx deckInfo: ", deckInfo)
    
    return (
        <div className="wrapper">
            <h1>Review Flashcards</h1>
            <FlashcardParent deckInfo={deckInfo} />
            <div id="detail">
                <Outlet />
            </div>
        </div>
    )
}

export default ReviewDeckRoute;
import React from 'react'
import { Outlet } from 'react-router-dom'
import axios from 'axios'
import Modal from '../modal'
import './deck.css'
import './main.css'
import DeckList from '../deck/deckList'

const DeckRoute = ({deckNames, setDeckNames, showModal, toggleShowModal, showEditOrReview, toggleShowEditOrReview, showAddDeck, toggleShowAddDeck, deckID, setDeckID, editDeckRoute, reviewDeckRoute, increaseReviewedCount}) => {
    const url = 'http://localhost:8080/api/decks'
    
    
    
    React.useEffect(() => {
        axios.get(url).then((response) => {
            setDeckNames(new Map(Object.entries(response.data)))
        })
    }, [setDeckNames])

    return (
        <>
            <div className='wrapper'>
                <div className='header'>
                    <h1>Your Flashcard Decks</h1>
                    <button onClick={ () => { 
                                                toggleShowModal()
                                                toggleShowAddDeck()
                                            }
                                    }>Add Deck</button>
                </div>
                <DeckList deckNameResponse={deckNames} toggleShowModal={toggleShowModal} setDeckID={setDeckID} toggleShowEditOrReview={toggleShowEditOrReview} />
                <Modal showModal={showModal} closeModal={toggleShowModal} showEditOrReview={showEditOrReview} toggleShowEditOrReview={toggleShowEditOrReview} showAddDeck={showAddDeck} toggleShowAddDeck={toggleShowAddDeck} onEditClick={editDeckRoute} onReviewClick={reviewDeckRoute} setID={deckID} increaseReviewedCount={increaseReviewedCount} />
            </div>
            <div id="detail">
                <Outlet />
            </div>
        </>
    )

}

export default DeckRoute;
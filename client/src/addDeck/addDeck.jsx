import React from 'react'
import axios from 'axios'

const AddDeck =({showAddDeck, toggleShowAddDeck}) => {

    const url = "http://localhost:8080/api/deck/"

    const [deckName, setDeckName] = React.useState("")

    const setInput = (e) => {
        setDeckName(e.target.value)
    }

    const submitForm = (e) => {
        e.preventDefault()
    }

    function clickSubmit() {
        axios.post(url, null, {params: {"deckName": deckName}})
    }

    if(!showAddDeck) {
        return null
    }

    return (
        <>
            <form onSubmit={ (e) => {submitForm(e)} } >
                <label>New Flashcard Deck Name</label>
                <input type="text" value={deckName} onChange={ (e) => {setInput(e)} } ></input>
            </form>
            <div className='submit'>
                <button onClick={clickSubmit} >Submit</button>
            </div>
        </>
    )
}

export default AddDeck;
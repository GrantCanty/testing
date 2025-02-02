import './deck.css'

function DeckInfo({id, name, toggleShowModal, setID, toggleShowEditOrReview}) {
    return (
        <div className="deck-item" id={`id${id%5}`} onClick={ () => { 
                                                                        setID(id)
                                                                        toggleShowModal()
                                                                        toggleShowEditOrReview()
                                                                    }
                                                                    } >
            <h2>{ name }</h2> 
        </div>
    )
}

export default DeckInfo;
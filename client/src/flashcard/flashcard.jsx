import React from 'react'
import './flashcard.css'

const Flashcard = ({deckInfo, cardPos, setCardPos, showTitle, toggleShowTitle}) => {

    if (deckInfo.length !== 0) {
        let card = new Map(Object.entries(deckInfo[cardPos]))

        return (
            <div className="flashcard" onClick={toggleShowTitle}>
                <div className='flashcard-info'>
                    { showTitle ? <h2>{card.get("topic")}</h2> : <h3>{card.get("description")}</h3> }
                    <button disabled={cardPos === 0 ? 'disabled' : null} className='new-card' onClick={(e) => {e.stopPropagation(); setCardPos(-1)}} >-</button>
                    <button disabled={cardPos === deckInfo.length - 1 ? 'disabled' : null } className='new-card' onClick={(e) => {e.stopPropagation(); setCardPos(1)}}  >+</button>
                    </div>
            </div>
        )
    }
}

export default Flashcard;
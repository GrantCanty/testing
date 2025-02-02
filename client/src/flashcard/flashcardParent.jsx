import React from "react";
import Flashcard from "./flashcard";


const FlashcardParent =({deckInfo}) => {
    const [showTitle, setShowTitle] = React.useState(true)

    function toggleShowTitle() {
        setShowTitle(!showTitle)

    }
    
    const deckLength = deckInfo.length

    const [cardPos, setCardPos] = React.useState(Number)

    function newCardPos(num) {
        if (num === -1 && cardPos > 0) {
            setCardPos(cardPos + num)
            setShowTitle(true)
        } 
        if (num === 1 && cardPos < deckLength -1) {
            setCardPos(cardPos + num)
            setShowTitle(true)
        }
    }

    return (
        <>
            <Flashcard deckInfo={deckInfo} cardPos={cardPos} setCardPos={newCardPos} showTitle={showTitle} toggleShowTitle={toggleShowTitle} />
        </>
    )

}

export default FlashcardParent;
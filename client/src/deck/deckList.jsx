import DeckInfo from "./deckInfo";

const DeckList = ({deckNameResponse, toggleShowModal, setDeckID, toggleShowEditOrReview}) => {
    return (
        <div className='deck-list'>
            { [...deckNameResponse.keys()].map((i) => {
                return <DeckInfo name={deckNameResponse.get(i)} key={i} id={i} toggleShowModal={toggleShowModal} setID={setDeckID} toggleShowEditOrReview={toggleShowEditOrReview} />
                })
            }
        </div>
    )
}

export default DeckList;
import "./card.css"

const EditFlashcard = ({flaschard, editState}) => {
    flaschard = new Map(Object.entries(flaschard))
    return (
            <div className="edit-flashcard" >
                <div className="edit-flashcard-topic">
                    <input type="text" name="topic" value={flaschard.get("topic")} onChange={ (event) => { editState(flaschard.get("id"), event) } } ></input>
                </div>
                <div className="edit-flashcard-desc">
                    <input type="text" name="description" value={flaschard.get("description")} onChange={ (event) => { editState(flaschard.get("id"), event) } } ></input>
                </div>
            </div>
    )
}

export default EditFlashcard;
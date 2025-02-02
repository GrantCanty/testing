import ReactDOM from 'react-dom'
import { useNavigate } from 'react-router-dom'
import AddDeck from './addDeck/addDeck'
import EditOrReview from './editOrReview/editOrReview'
import './modal.css'

const Modal = ({showModal, closeModal, showEditOrReview, toggleShowEditOrReview, showAddDeck, toggleShowAddDeck, onEditClick, onReviewClick, setID, increaseReviewedCount}) => {
    const navigate = useNavigate()

    if (!showModal) {
        return null
    }

    return ReactDOM.createPortal (
        <div className="modal-bg">
            <div className="modal">
                < EditOrReview onEditClick={onEditClick} onReviewClick={onReviewClick} setID={setID} increaseReviewedCount={increaseReviewedCount} closeModal={closeModal} showEditOrReview={showEditOrReview} toggleShowEditOrReview={toggleShowEditOrReview} navigate={navigate} />
                < AddDeck showAddDeck={showAddDeck} toggleShowAddDeck={toggleShowAddDeck} />
                <button onClick={ () => { 
                                            closeModal()
                                            if(showEditOrReview === true) {toggleShowEditOrReview()}
                                            if(showAddDeck === true) {toggleShowAddDeck()}
                                        }
                                }>Close Modal</button>
            </div>
        </div>, document.body
    )
}

export default Modal;
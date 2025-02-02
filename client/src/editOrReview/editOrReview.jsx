const EditOrReview = ({onEditClick, onReviewClick, setID, increaseReviewedCount, closeModal, showEditOrReview, toggleShowEditOrReview, navigate}) => {
    
    if(!showEditOrReview) {
        return null
    }

    return (
        <div className='options-list'>
            <div className='options-item edit' onClick={() => {
                                                                navigate(onEditClick + '/' + setID)
                                                                closeModal()
                                                                toggleShowEditOrReview()
                                                                }} >
                <h3>Edit Deck</h3>
            </div>
            <div className='options-item review' onClick={() => {
                                                                navigate(onReviewClick + '/' + setID)
                                                                increaseReviewedCount(1)
                                                                closeModal()
                                                                toggleShowEditOrReview()
                                                                }} >
                <h3>Review Deck</h3>
            </div>
        </div>
    )
}

export default EditOrReview;
import React from 'react'

const useEditOrReview = () => {
    const [isShowing, setIsShowing] = React.useState(false)

    function toggle() {
        setIsShowing(!isShowing)
    }

    return [
        isShowing,
        toggle
    ];
}

export default useEditOrReview
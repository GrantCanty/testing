import React from 'react'

const useAddDeck = () => {
    const [isShowing, setIsShowing] = React.useState(false)

    function toggle() {
        setIsShowing(!isShowing)
    }

    return [
        isShowing,
        toggle
    ];
}

export default useAddDeck
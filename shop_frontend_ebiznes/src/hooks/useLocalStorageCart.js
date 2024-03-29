import {useState} from "react";

const getSavedValue = () => {
    const savedValue = JSON.parse(localStorage.getItem("spaniel_shopping_cart"));
    if (savedValue) return savedValue;
    return [];
}

const useLocalStorageCart = () => {
    const [cartItems, setStoredValue] = useState(() => {
        return getSavedValue();
    })

    const setCartState = (value) => {
        setStoredValue(value);
        localStorage.setItem("spaniel_shopping_cart", JSON.stringify(value))
    }

    const addItemToCart = (item) => {
        setStoredValue(oldCartItems => {
            if (oldCartItems) {
                const newCart = [...oldCartItems, item];
                localStorage.setItem("spaniel_shopping_cart", JSON.stringify(newCart))
                return [...oldCartItems, item]
            } else {
                localStorage.setItem("spaniel_shopping_cart", JSON.stringify([item]))
                return [item]
            }
        })
    }

    const cleanBasket = () => {
        setStoredValue([]);
        localStorage.setItem("spaniel_shopping_cart", null)
    }

    return {cartItems, addItemToCart, setCartState, cleanBasket};
}

export default useLocalStorageCart;
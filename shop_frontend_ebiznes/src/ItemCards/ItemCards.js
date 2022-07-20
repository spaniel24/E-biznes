import ItemCard from "./ItemCard/ItemCard";
import './ItemCards.css';
import {useCallback, useEffect, useState} from "react";
import axios from "axios";
import useLocalStorageCart from "../hooks/useLocalStorageCart";

const ItemCards = ({category}) => {

    const [items, setItems] = useState([]);
    const {addItemToCart} = useLocalStorageCart();

    const fetchItems = useCallback(
        () => {
            axios.get(`http://localhost:8080/products?category=${category}`).then(res => {
                setItems(res.data)
            })
        }, [category]);

    useEffect(() => {
        fetchItems();
    }, [fetchItems])

    const handleAddToCart = (item) => {
        addItemToCart(item);
    }

    const renderItems = () => {
        if (items) {
            const itemCards = [];
            items.forEach(item => {
                itemCards.push(<ItemCard key={item.ID} item={item} clickHandle={handleAddToCart}/>)
            })
            return itemCards;
        }
    }

    return (
        <div className="item-cards">
            {renderItems()}
        </div>
    );
}

export default ItemCards;
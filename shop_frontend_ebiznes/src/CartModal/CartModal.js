import './CartModal.css'
import useLocalStorageCart from "../hooks/useLocalStorageCart";
import axios from "axios";

const CartModal = ({onClose, setShowingPaymentModal}) => {
    // return ReactDOM.createPortal(

    const {cartItems} = useLocalStorageCart();

    const getCartItemsList = () => {
        let id = 0;
        return cartItems.map(cartItemsObject => {
                id += 1;
                return (
                    <li key={id}>
                        <span>{cartItemsObject.Name}, {cartItemsObject.Price} zł</span>
                    </li>
                )
            }
        );
    };

    const getCartTotalPrice = () => {
        let totalPrice = 0;
        cartItems.forEach(cartItem => {
            totalPrice = totalPrice + cartItem.Price;
        })
        return totalPrice;
    }

    const onPay = () => {
        const urlParams = new URLSearchParams(window.location.search);
        axios.post(`https://shopworkingbackend.azurewebsites.net/order?user_token=${urlParams.get('user_token')}`, {
            Price: getCartTotalPrice(),
            OrderProducts: getCartItemsList(),
            Status: 'Ready to pay'
        }).then((res) => {
            console.log(res)
            setShowingPaymentModal(true);
            onClose();
        }).catch(()=>{
            alert("You need to login before making order")
        })
    }

    return (
        <>
            <div className="transparent-background"/>
            <div className="cart-modal">
                <div>
                    In cart:
                    <ul>
                        {getCartItemsList()}
                    </ul>
                </div>
                <div>
                    To pay: {getCartTotalPrice()}
                </div>
                <div className="cart-modal-buttons">
                    <button onClick={onClose}>Cancel</button>
                    <button onClick={onPay}>Go to Payment</button>
                </div>
            </div>
        </>
    )
}

export default CartModal;
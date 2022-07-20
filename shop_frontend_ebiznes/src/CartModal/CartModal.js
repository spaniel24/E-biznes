import './CartModal.css'
import useLocalStorageCart from "../hooks/useLocalStorageCart";

const CartModal = ({onClose, setShowingPaymentModal}) => {
    // return ReactDOM.createPortal(

    const {cartItems} = useLocalStorageCart();

    const getCartItemsList = () => {
        let id = 0;
        return cartItems.map(cartItemsObject => {
                id += 1;
                return (
                    <li key={id}>
                        <span>{cartItemsObject.name}, {cartItemsObject.price} zł</span>
                    </li>
                )
            }
        );
    };

    const getCartTotalPrice = () => {
        let totalPrice = 0;
        cartItems.forEach(cartItem => {
            totalPrice = totalPrice + cartItem.price;
        })
        return totalPrice;
    }

    const onPay = () => {
        setShowingPaymentModal(true);
        onClose();
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
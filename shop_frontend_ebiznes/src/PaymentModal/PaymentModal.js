import './PaymentModal.css'
import {useState} from "react";
import axios from "axios";
import useLocalStorageCart from "../hooks/useLocalStorageCart";

const PaymentModal = ({onClose}) => {
    const [creditCardNumber, setCreditCardNumber] = useState('');
    const [waiting, setWaiting] = useState(false);
    const {cleanBasket} = useLocalStorageCart();

    const handleSubmit = () => {
        setWaiting(true);
        const urlParams = new URLSearchParams(window.location.search);
        axios.post(`http://localhost:8080/payments?user_token=${urlParams.get('user_token')}`).then(() => {
            alert('Payment has ended successfully, come again!')
            setWaiting(false);
            cleanBasket();
            onClose();
        })
    };

    const renderWaiting = () => {
        if (waiting) {
            return(
                <div>
                    Processing payment...
                </div>
            );
        }
    }
    return (
        <div className="payment-modal">
            <label>Card Number:</label>
            <textarea
                id="cardNumber"
                name="cardNumber"
                onChange={(event) => setCreditCardNumber(event.target.value)}
                value={creditCardNumber}
            />
            <button onClick={handleSubmit}>Send money</button>
            {renderWaiting()}
        </div>
    )
}

export default PaymentModal;
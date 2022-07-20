import './PaymentModal.css'
import {useState} from "react";
import axios from "axios";

const PaymentModal = ({onClose}) => {
    const [creditCardNumber, setCreditCardNumber] = useState('');
    const [waiting, setWaiting] = useState(false);

    const handleSubmit = () => {
        setWaiting(true);
        axios.post('http://localhost:8080/payments').then(() => {
            alert('Payment has ended successfully, come again!')
            setWaiting(false);
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
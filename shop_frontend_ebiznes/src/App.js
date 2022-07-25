import Navbar from "./Navbar/Navbar";
import LeftSideMenu from "./LeftSideMenu/LeftSideMenu";
import './App.css'
import ItemCards from "./ItemCards/ItemCards";
import {useCallback, useEffect, useState} from "react";
import axios from 'axios'
import CartModal from "./CartModal/CartModal";
import PaymentModal from "./PaymentModal/PaymentModal";
import LoginModal from "./LoginModal/LoginModal";

function App() {

    const [categories, setCategories] = useState([]);
    const [selectedCategory, setSelectedCategory] = useState('');
    const [showingOpenModal, setShowingOpenModal] = useState(false);
    const [showingPaymentModal, setShowingPaymentModal] = useState(false);
    const [showingLoginModal, setShowingLoginModal] = useState(false);

    const fetchCategories = useCallback(
        () => {
            axios.get('http://shopworkingbackend.azurewebsites.net/categories').then(res => {
                console.log(res.data)
                setCategories(res.data)
            })
        }, []);

    useEffect(() => {
        fetchCategories();
    }, [fetchCategories])

    const renderCartModal = () => {
        if (showingOpenModal) {
            return (
                <CartModal onClose={()=>setShowingOpenModal(false)} setShowingPaymentModal={setShowingPaymentModal}/>
            )
        }

        return null;
    }

    const renderPaymentModal = () => {
        if (showingPaymentModal) {
            return (
                <PaymentModal onClose={()=>setShowingPaymentModal(false)}/>
            )
        }

        return null;
    }

    const renderLoginModal = () => {
        if (showingLoginModal) {
            return (
                <LoginModal onClose={()=>setShowingLoginModal(false)}/>
            )
        }

        return null;
    }

    return (
        <>
            {renderCartModal()}
            {renderPaymentModal()}
            {renderLoginModal()}
            <div className="website">
                <div className="website-top">
                    <Navbar showCartModal={() => setShowingOpenModal(true)} showLoginModal={()=>setShowingLoginModal(true)}/>
                </div>
                <div className="website-middle">
                    <LeftSideMenu categories={categories} categoryCallback={category => setSelectedCategory(category)}/>
                    <ItemCards category={selectedCategory}/>
                </div>
            </div>
        </>
    );
}

export default App;

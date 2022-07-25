import './Navbar.css';
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome"
import {faShoppingCart} from "@fortawesome/free-solid-svg-icons";

const Navbar = ({showCartModal, showLoginModal}) => {

    return (
        <div className="navbar">
            <div className="left-navbar">
            </div>
            <div className="title">
                Spaniel-shop
            </div>
            <div>
                <button className="cart" onClick={showLoginModal}>
                    Login
                </button>
                <button className="cart" onClick={showCartModal}>
                    <FontAwesomeIcon icon={faShoppingCart}/>
                </button>
            </div>
        </div>
    );
}
export default Navbar;
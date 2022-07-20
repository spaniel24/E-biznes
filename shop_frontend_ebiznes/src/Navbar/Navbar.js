import './Navbar.css';
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome"
import {faShoppingCart} from "@fortawesome/free-solid-svg-icons";

const Navbar = ({showCartModal}) => {
    return (
        <div className="navbar">
            <div className="left-navbar">
            </div>
            <div className="title">
                Spaniel-shop
            </div>
            <button className="cart" onClick={showCartModal}>
                <FontAwesomeIcon icon={faShoppingCart}/>
            </button>
        </div>
    );
}
export default Navbar;
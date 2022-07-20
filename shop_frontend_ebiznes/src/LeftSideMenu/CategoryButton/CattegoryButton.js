import {FontAwesomeIcon} from "@fortawesome/react-fontawesome"
import {faAngleRight} from "@fortawesome/free-solid-svg-icons";
import './CategoryButton.css';

const CategoryButton = ({categoryName, categoryCallback}) => {
    return (

        <button className="no-style-button" onClick={() => categoryCallback(categoryName)}>
            <div className="category-button">
                <span>{categoryName}</span>
                <FontAwesomeIcon icon={faAngleRight}/>
            </div>
        </button>
    )
}

export default CategoryButton;

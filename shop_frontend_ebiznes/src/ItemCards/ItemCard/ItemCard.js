import './ItemCard.css'

const ItemCard = ({item, clickHandle}) => {

    return (
        <div className="item-card" key={item._id}>
            <img src={item.ImageUrl}
                 alt={item.Name}/>
            <div className="description">
                <div>
                    {item.Name}
                </div>
                <div>
                    {item.Description}
                </div>
                <div>
                    {item.Category}
                </div>
            </div>
            <div className="right-side">
                <div>
                    {item.Price} zł
                </div>
                <button className="add-to-cart-button" type="button" onClick={() => clickHandle(item)}>
                    Add to cart
                </button>
            </div>
        </div>
    )
};
export default ItemCard;
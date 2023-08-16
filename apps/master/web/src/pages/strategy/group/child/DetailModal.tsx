import React from "react";
import type {GroupItem} from "@/apis/prom/prom";
import {Modal} from "@arco-design/web-react";
import groupStyle from "../style/group.module.less";
import StrategyModal from "@/pages/strategy/child/StrategyModal";
import StrategyList from "@/pages/strategy/group/child/StrategyList";

export interface DetailModalProps {
    item?: GroupItem
    children?: React.ReactNode
}



const DetailModal: React.FC<DetailModalProps> = (props) => {
    const {item, children} = props;
    const [visible, setVisible] = React.useState(false);

    const handleOpenModal = () => {
        setVisible(true);
    }

    const handleCloseModal = () => {
        setVisible(false);
    }

    return (
        <>
            <div onClick={handleOpenModal}>{children}</div>
            <Modal
                visible={visible}
                onCancel={handleCloseModal}
                escToExit
                footer={null}
                className={groupStyle.GroupDetail}
                style={{width: "100%"}}
                title={item?.name}
                unmountOnExit
            >
                <StrategyList groupItem={item} />
            </Modal>
        </>
    )
}

export default DetailModal;
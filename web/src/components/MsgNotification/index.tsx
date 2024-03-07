import { getWsURL } from '@/apis/request'
import { Message } from '@/apis/types'
import { createWebSocket } from '@/utils/ws'
import { Statistic, Typography, notification } from 'antd'
import dayjs from 'dayjs'
import React, { useEffect, useState } from 'react'

export interface MsgProps {}

const { Countdown } = Statistic
const { Paragraph } = Typography

export const MsgNotification: React.FC<MsgProps> = (props) => {
    const {} = props

    const [api, contextHolder] = notification.useNotification()

    const openNotification = (msg: Message) => {
        const { content, title } = msg
        const duration = 3

        const nowTime = dayjs().add(duration, 's').format('YYYY-MM-DD HH:mm:ss')

        api.info({
            message: title,
            description: (
                <div>
                    <Countdown
                        style={{ position: 'absolute', top: 8, right: 8 }}
                        valueStyle={{ fontSize: 20 }}
                        value={nowTime}
                        format="ss"
                        suffix="s"
                    />
                    <Paragraph
                        ellipsis={{ rows: 2, expandable: true, symbol: 'more' }}
                    >
                        {content}
                    </Paragraph>
                </div>
            ),
            style: { position: 'relative' },
            closeIcon: null,
            placement: 'topRight',
            duration: duration
        })
    }

    const [ws] = useState(createWebSocket(getWsURL()))
    useEffect(() => {
        ws.onmessage = (event) => {
            const msg = JSON.parse(event.data)
            openNotification(msg)
        }
    }, [ws])
    return <>{contextHolder}</>
}

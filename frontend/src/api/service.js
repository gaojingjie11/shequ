import request from '@/utils/request'

// 获取公告列表
export function getNoticeList() {
    return request({
        url: '/notices',
        method: 'get'
    })
}

// 获取公告详情
export function getNoticeDetail(id) {
    return request({
        url: `/notice/${id}`,
        method: 'get'
    })
}

// 标记公告已读
export function readNotice(id) {
    return request({
        url: `/notice/read/${id}`,
        method: 'post'
    })
}

// 创建报修
export function createRepair(data) {
    return request({
        url: '/repair/create',
        method: 'post',
        data
    })
}

// 获取报修列表
export function getRepairList() {
    return request({
        url: '/repair/list',
        method: 'get'
    })
}

// 创建访客登记
export function createVisitor(data) {
    return request({
        url: '/visitor/create',
        method: 'post',
        data
    })
}

// 获取访客列表
export function getVisitorList() {
    return request({
        url: '/visitor/list',
        method: 'get'
    })
}

// 获取我的车位
export function getMyParking() {
    return request({
        url: '/parking/my',
        method: 'get'
    })
}

// 绑定车牌
export function bindCar(data) {
    return request({
        url: '/parking/bind',
        method: 'post',
        data
    })
}

// 获取物业费列表
export function getPropertyFeeList() {
    return request({
        url: '/property/list',
        method: 'get'
    })
}

// 缴纳物业费
export const payPropertyFee = (data) => {
    return request({
        url: '/finance/pay',
        method: 'post',
        data: {
            business_id: data.related_id,
            pay_type: 2  // 2表示物业费
        }
    })
}

// 获取门店列表
export function getStoreList() {
    return request({
        url: '/stores',
        method: 'get'
    })
}

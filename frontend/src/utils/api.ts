import Axios from 'axios'
import { urlParser } from './urlParser'
const token = new URLSearchParams(window.location.search).get('token')
export const axios = Axios.create({
	baseURL: urlParser('HTTP', 'api'),
	// baseURL: 'http://10.0.222.181:3000/api',
	headers: {
		common: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	},
})

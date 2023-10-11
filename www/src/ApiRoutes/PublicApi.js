import ApiRequestMethod from './helper';

const ApiRequest = async (route, body, methodType) => {
	const url = `${ApiRequestMethod.url}/public/${route}`;
	const request = ApiRequestMethod[methodType];
	const bodyOfRequest = JSON.stringify(body);

	const jsonBodyConfig = {
		headers: {
			'Content-Type': 'application/json',
			Accept: 'application/json',
		},
	};
	const response = !(methodType === 'get' || methodType === 'delete')
		? await request(url, bodyOfRequest, jsonBodyConfig)
		: await request(url);

	return response;
};
export default ApiRequest;

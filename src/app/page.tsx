import {
	Table,
} from "react-bootstrap";



const fetchData = async () => {
	const response = await fetch("http://localhost:3000/today");
	return await response.json();
};

export default async function Home() {
	const data = await fetchData();

	return (
		<div>
			<h1>오늘의 학식 메뉴</h1>
			<Table>
				<thead>
					<tr>
						<th>아침</th>
						<th>점심</th>
						<th>저녁</th>
					</tr>
				</thead>
				<tbody>
					<tr key={data.id}>
						<td>{data.table.breakfast}</td>
						<td>{data.table.lunch}</td>
						<td>{data.table.dinner}</td>
					</tr>
				</tbody>
			</Table>
		</div>
	);
};
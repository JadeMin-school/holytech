'use client';

import { useState } from "react";

import {
	Button, Table,
} from "react-bootstrap";



const fetchData = async () => {
	const response = await fetch("https://jsonplaceholder.typicode.com/posts");
	return await response.json();
};

export default function Home() {
	const [data, setData] = useState<any>(null);

	return (
		<div>
			<h1>Home</h1>
			<Button
				className=""
				onClick={
					async () => {
						setData(await fetchData())
					}
				}>
					Fetch Data
				</Button>
			<Table>
				<thead>
					<tr>
						<th>날짜</th>
						<th>메뉴</th>
					</tr>
				</thead>
				<tbody>
					{data && data.map((item: any) => (
						<tr key={item.id}>
							<td>{item.userId}</td>
							<td>{item.id}</td>
						</tr>
					))}
				</tbody>
			</Table>
		</div>
	);
};
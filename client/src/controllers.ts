import { PORT } from "./utils";

//* Expenses/Incomes controllers
export async function GetTransactions(table: "incomes" | "expenses" | null) {
  if (!table) return
  return new Promise(async (resolve, reject) => {
    const response = await fetch(`http://localhost:${PORT}/api/transactions/${table}`, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();
    if (data.status == "error") resolve(null);
    if (data.status == "success") resolve(data.response);
  });
}

export async function EditTransaction(info: Transaction | null) {
  if (!info) return
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify({
      ...info,
      amount: parseFloat(String(info.amount)),
      fulfilled: !!info.fulfilled,
      tax: parseFloat(String(info.tax)),
      taxxed: !!info.taxxed,
      recurring: !!info.recurring,
    });
    const response = await fetch(`http://localhost:${PORT}/api/transactions/${info.table}`, {
      method: "PATCH",
      body: body,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res) => res.json());
    console.log(response.response);
    resolve(null);
  });
}

export async function AddTransaction(info: Transaction | null) {
  if (!info) return
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify({
      ...info,
      amount: parseFloat(String(info.amount)),
      fulfilled: !!info.fulfilled,
      tax: parseFloat(String(info.tax)),
      taxxed: !!info.taxxed,
      recurring: !!info.recurring,
    });

    const response = await fetch(`http://localhost:${PORT}/api/transactions/${info.table}`, {
      method: "POST",
      body: body,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res) => res.json());
    console.log(response.response);
    resolve(null);
  });
}

export async function DeleteTransaction(info: Transaction | null) {
  if (!info) return
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify({ Id: info.id });
    const response = await fetch(`http://localhost:${PORT}/api/transactions/${info.table}`, {
      method: "DELETE",
      body: body,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res) => res.json());
    console.log(response.response);
    resolve(null);
  });
}

//* Categories controllers
export async function deleteCategory(info: Category) {
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify({ Id: info.id });
    const response = await fetch(`http://localhost:${PORT}/api/categories`, {
      method: "DELETE",
      body: body,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res) => res.json());
    console.log(response.response);
    resolve(null);
  });
}

export async function addCategory(info: Category) {
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify({
      ...info,
      budget: parseFloat(String(info.budget)),
      spent: parseFloat(String(info.spent)),
      gotten: parseFloat(String(info.gotten)),
    });
    const response = await fetch(`http://localhost:${PORT}/api/categories`, {
      method: "POST",
      body: body,
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    }).then((res) => res.json());
    console.log(response.response);
    resolve(null);
  });
}

export async function editCategory(info: Category | null) {
  if (!info) return
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify({
      ...info,
      spent: parseFloat(String(info.spent)),
      budget: parseFloat(String(info.budget)),
      gotten: parseFloat(String(info.gotten)),
    });
    const response = await fetch(`http://localhost:${PORT}/api/categories`, {
      method: "PATCH",
      body: body,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res: Response) => res.json());
    console.log(response.response);
    resolve(null);
  });
}
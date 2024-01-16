interface Transaction {
  table: "incomes" | "expenses" | null;
  id: number | null;
  parentid: number | null;
  title: string | null;
  description: string | null;
  currency: string | null;
  amount: number | null;
  paymentmethod: string | null;
  participant: string | null;
  recurring: boolean | null;
  interval: string | null;
  category: string | null;
  taxxed: boolean | null;
  tax: number | null;
  fulfilled: boolean | null;
  datecreated: string | null;
  fileurl: string | null;
}
interface Category {
  id: number | null;
  parentid: number | null;
  title: string | null;
  description: string | null;
  currency: string | null;
  budget: number | null;
  spent: number | null;
  gotten: number | null;
  type: string | null;
  budgeted: boolean | null;
}
interface User {
  username: string | null;
  email: string | null;
  password: string | null;
  currency: string | null;
  balance: number | null;
}
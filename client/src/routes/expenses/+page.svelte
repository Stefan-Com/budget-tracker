<script lang="ts">
  import { onMount } from "svelte";
  import {
    checkIfValidNumberInput,
    CURRENCIES,
    getCategories,
    paymentMethods,
  } from "../../utils";
  import ExpenseForm from "../../components/forms/ExpenseForm.svelte";
  import EditBtn from "../../components/buttons/EditBtn.svelte";
  import DeleteBtn from "../../components/buttons/DeleteBtn.svelte";
  import InfoBtn from "../../components/buttons/InfoBtn.svelte";
  import {
    AddTransaction,
    DeleteTransaction,
    EditTransaction,
    GetTransactions,
  } from "../../controllers";
  let form: any;

  let currentExpense: Transaction | null;
  let currentEditExpense: Transaction | null;
  let expenses: any | null = [];
  let categories: any | null = [];

  onMount(async () => {
    categories = await getCategories();
    categories = categories.filter(
      (category: Category) => category.type == "expense",
    );
    if (!categories) categories = [];
    expenses = await GetTransactions("expenses");
    if (!expenses) expenses = [];
    form = document.getElementById("form") || { hidden: false };
  });

  function handleDisplayInfo(info: Transaction) {
    currentExpense = info;
  }
  function handleEditInfo(info: Transaction) {
    currentEditExpense = info;
  }
</script>

<div>
  <h1 class="pl-5">Your Expenses:</h1>
  <ul>
    {#each expenses as expense}
      <li class="hover:bg-gray-200 flex justify-between w-1/2 p-5 rounded-3xl">
        <div>
          Name: {expense.title}
          <br />
          Amount: {expense.amount}
        </div>
        <div class="flex justify-around items-center w-28">
          <EditBtn
            handleClick={() => {
              handleEditInfo(expense);
            }}
          />
          <DeleteBtn
            handleClick={async () => {
              await DeleteTransaction(expense);
              expenses = await GetTransactions("expenses");
              if (!expenses) expenses = [];
            }}
          />
          <InfoBtn
            handleClick={() => {
              handleDisplayInfo(expense);
            }}
          />
        </div>
      </li>
    {/each}
    <li
      class="hover:bg-stone-600 flex justify-between w-1/2 p-5 rounded-3xl bg-stone-400"
    >
      <button
        class="w-full"
        on:click={() => {
          form.hidden = !form.hidden;
        }}>Add New Expense</button
      >
    </li>
  </ul>
  {#if currentExpense}
    <dialog
      data-te-toggle="modal"
      id="modal"
      open
      class="border-4 border-black top-0 left-0 w-96 h-fit"
    >
      <div>
        <h3>Title: {currentExpense.title}</h3>
      </div>
      <br />
      <div>
        <h3>Description: {currentExpense.description}</h3>
      </div>
      <br />
      <div>
        <h3>Currency: {currentExpense.currency}</h3>
      </div>
      <br />
      <div>
        <h3>Amount: {currentExpense.amount}</h3>
      </div>
      <br />
      <div>Payment Method: {currentExpense.paymentmethod}</div>
      <br />
      <div>
        <h3>
          Merchent(the one who's getting paid):{currentExpense.participant}
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Recurring: {!!currentExpense.recurring ? "Yes" : "No"}
        </h3>
        {#if !!currentExpense.recurring}
          <h3>Interval: {currentExpense.interval}</h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>Category: {currentExpense.category}</h3>
      </div>
      <br />
      <div>
        <h3>Taxxed: {!!currentExpense.tax ? "Yes" : "No"}</h3>
        {#if !!currentExpense.tax}
          <h3>Tax: {currentExpense.tax}</h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>Paid: {!!currentExpense.fulfilled ? "Yes" : "No"}</h3>
      </div>
      <br />
      <div>
        <h3>File URL: {currentExpense.fileurl}</h3>
      </div>
      <br />
      <div>
        <button
          on:click={() => {
            currentExpense = null;
          }}>Exit</button
        >
      </div>
      <br />
    </dialog>
  {:else if currentEditExpense}
    <dialog
      data-te-toggle="modal"
      id="modal"
      open
      class="border-4 border-black top-0 left-0 w-96 h-fit"
    >
      <div>
        <h3>Title: <input bind:value={currentEditExpense.title} /></h3>
      </div>
      <br />
      <div>
        <h3>
          Description: <input bind:value={currentEditExpense.description} />
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Currency: <select bind:value={currentEditExpense.currency}>
            {#each CURRENCIES as currency}
              <option value={currency}>{currency}</option>
            {/each}
          </select>
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Amount: <input
            bind:value={currentEditExpense.amount}
            on:input={checkIfValidNumberInput}
          />
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Payment Method: <select bind:value={currentEditExpense.paymentmethod}>
            {#each paymentMethods as method}
              <option value={method}>{method}</option>
            {/each}
          </select>
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Merchent: <input bind:value={currentEditExpense.participant} />
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Is recurring: <input
            type="checkbox"
            bind:checked={currentEditExpense.recurring}
          />
        </h3>
        {#if currentEditExpense.recurring}
          <h3>
            Interval: <input
              bind:value={currentEditExpense.interval}
              on:input={checkIfValidNumberInput}
            />
          </h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>
          Category: <select bind:value={currentEditExpense.category}>
            {#each categories as category}
              <option selected value={category.title}>{category.title}</option>
            {/each}
          </select>
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Taxxed: <input
            type="checkbox"
            bind:checked={currentEditExpense.taxxed}
          />
        </h3>
        {#if currentEditExpense.taxxed}
          <h3>
            Tax: <input
              bind:value={currentEditExpense.tax}
              on:input={checkIfValidNumberInput}
            />
          </h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>
          Paid: <input
            type="checkbox"
            bind:checked={currentEditExpense.fulfilled}
          />
        </h3>
      </div>
      <br />
      <div>
        <button
          on:click={async () => {
            await EditTransaction(currentEditExpense);
            currentEditExpense = null;
            expenses = await GetTransactions("expenses");
            if (!expenses) expenses = [];
          }}>Save changes</button
        >
        <button
          on:click={() => {
            currentEditExpense = null;
          }}>Exit</button
        >
      </div>
      <br />
    </dialog>
  {/if}
  <ExpenseForm
    {categories}
    addExpense={AddTransaction}
    getExpenses={() => {
      return new Promise(async (resolve, reject) => {
        expenses = await GetTransactions("expenses");
        if (!expenses) expenses = [];
        resolve(null);
      });
    }}
  />
</div>

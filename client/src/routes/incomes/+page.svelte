<script lang="ts">
  import { onMount } from "svelte";
  import {
    checkIfValidNumberInput,
    CURRENCIES,
    getCategories,
    paymentMethods,
  } from "../../utils";
  import IncomeForm from "../../components/forms/IncomeForm.svelte";
  import EditBtn from "../../components/buttons/EditBtn.svelte";
  import DeleteBtn from "../../components/buttons/DeleteBtn.svelte";
  import InfoBtn from "../../components/buttons/InfoBtn.svelte";
  import {
    AddTransaction,
    DeleteTransaction,
    EditTransaction,
    GetTransactions,
  } from "../../controllers";
  let showIncomes = false;

  let form: any;
  let categories: any = [];
  let currentIncome: Transaction | null;
  let currentEditIncome: Transaction | null;
  let incomes: any | null;

  onMount(async () => {
    categories = await getCategories();
    if (!categories) categories = [];
    incomes = await GetTransactions("incomes");
    if (!incomes) incomes = [];
    showIncomes = true;
    form = document.getElementById("form") || { hidden: false };
  });

  function displayInfo(info: Transaction) {
    currentIncome = info;
  }
  function editInfo(info: Transaction) {
    currentEditIncome = info;
  }
</script>

<div>
  <h1 class="pl-5">Your Incomes:</h1>
  <ul>
    {#if showIncomes}
      {#each incomes as income}
        <li
          class="hover:bg-gray-200 flex justify-between w-1/2 p-5 rounded-3xl"
        >
          <div>
            Name: {income.title}
            <br />
            Amount: {income.amount}
          </div>
          <div class="flex justify-around items-center w-28">
            <EditBtn
              handleClick={() => {
                editInfo(income);
              }}
            />
            <DeleteBtn
              handleClick={async () => {
                await DeleteTransaction(income);
                incomes = await GetTransactions("incomes");
                if (!incomes) incomes = [];
              }}
            />
            <InfoBtn
              handleClick={() => {
                displayInfo(income);
              }}
            />
          </div>
        </li>
      {/each}
    {/if}
    <li
      class="hover:bg-stone-600 flex justify-between w-1/2 p-5 rounded-3xl bg-stone-400"
    >
      <button
        class="w-full"
        on:click={() => {
          form.hidden = !form.hidden;
        }}>Add New Income</button
      >
    </li>
  </ul>
  {#if currentIncome}
    <dialog
      data-te-toggle="modal"
      id="modal"
      open
      class="border-4 border-black top-0 left-0 w-96 h-fit"
    >
      <div>
        <h3>Title: {currentIncome.title}</h3>
      </div>
      <br />
      <div>
        <h3>Description: {currentIncome.description}</h3>
      </div>
      <br />
      <div>
        <h3>Currency: {currentIncome.currency}</h3>
      </div>
      <br />
      <div>
        <h3>Amount: {currentIncome.amount}</h3>
      </div>
      <br />
      <div>
        <h3>Payment Method: {currentIncome.paymentmethod}</h3>
      </div>
      <br />
      <div>
        <h3>Client(the one who's paying): {currentIncome.participant}</h3>
      </div>
      <br />
      <div>
        <h3>
          Is recurring: {!!currentIncome.interval ? "Yes" : "No"}
        </h3>
        {#if !!currentIncome.interval}
          <h3>Interval: {currentIncome.interval}</h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>Category: {currentIncome.category}</h3>
      </div>
      <br />
      <div>
        <h3>Taxxed: {!!currentIncome.tax ? "Yes" : "No"}</h3>
        {#if !!currentIncome.tax}
          <h3>Tax: {currentIncome.tax}</h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>Received: {!!currentIncome.fulfilled ? "Yes" : "No"}</h3>
      </div>
      <br />
      <div>
        <h3>File URL: {currentIncome.fileurl}</h3>
      </div>
      <br />
      <div>
        <button
          on:click={() => {
            currentIncome = null;
          }}>Exit</button
        >
      </div>
      <br />
    </dialog>
  {:else if currentEditIncome}
    <dialog
      data-te-toggle="modal"
      id="modal"
      open
      class="border-4 border-black top-0 left-0 w-96 h-fit"
    >
      <div>
        <h3>Title: <input bind:value={currentEditIncome.title} /></h3>
      </div>
      <br />
      <div>
        <h3>
          Description: <input bind:value={currentEditIncome.description} />
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Currency: <select bind:value={currentEditIncome.currency}>
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
            bind:value={currentEditIncome.amount}
            on:input={checkIfValidNumberInput}
          />
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Payment Method: <select bind:value={currentEditIncome.paymentmethod}>
            {#each paymentMethods as method}
              <option value={method}>{method}</option>
            {/each}
          </select>
        </h3>
      </div>
      <br />
      <div>
        <h3>Client: <input bind:value={currentEditIncome.participant} /></h3>
      </div>
      <br />
      <div>
        <h3>
          Is recurring: <input
            type="checkbox"
            bind:value={currentEditIncome.recurring}
          />
        </h3>
        {#if currentEditIncome.recurring}
          <h3>
            Interval: <input
              bind:value={currentEditIncome.interval}
              on:input={checkIfValidNumberInput}
            />
          </h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>
          Category: <select bind:value={currentEditIncome.category}>
            {#each categories as category}
              <option value={category.title}>{category.title}</option>
            {/each}
          </select>
        </h3>
      </div>
      <br />
      <div>
        <h3>
          Taxxed: <input
            type="checkbox"
            bind:checked={currentEditIncome.taxxed}
          />
        </h3>
        {#if currentEditIncome.taxxed}
          <h3>
            Tax: <input
              bind:value={currentEditIncome.tax}
              on:input={checkIfValidNumberInput}
            />
          </h3>
        {/if}
      </div>
      <br />
      <div>
        <h3>
          Received: <input
            type="checkbox"
            bind:checked={currentEditIncome.fulfilled}
          />
        </h3>
      </div>
      <br />
      <div>
        <button
          on:click={async () => {
            await EditTransaction(currentEditIncome);
            currentEditIncome = null;
            incomes = await GetTransactions("incomes");
            if (!incomes) incomes = [];
          }}>Save changes</button
        >
        <button
          on:click={() => {
            currentEditIncome = null;
          }}>Exit</button
        >
      </div>
      <br />
    </dialog>
  {/if}
  <IncomeForm
    {categories}
    addIncome={AddTransaction}
    getIncomes={() => {
      return new Promise(async (resolve, reject) => {
        incomes = await GetTransactions("incomes");
        if (!incomes) incomes = [];
        resolve(null);
      });
    }}
  />
</div>

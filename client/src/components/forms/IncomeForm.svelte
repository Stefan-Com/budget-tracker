<script lang="ts">
  import {
    checkIfValidNumberInput,
    CURRENCIES,
    paymentMethods,
  } from "../../utils";
  const info: Transaction = {
    table: "incomes",
    id: null,
    parentid: null,
    title: "",
    description: "",
    currency: CURRENCIES[0],
    amount: 0,
    paymentmethod: paymentMethods[0],
    participant: "",
    recurring: false,
    interval: "",
    category: "",
    taxxed: false,
    tax: 0,
    fulfilled: true,
    datecreated: null,
    fileurl: "",
  };
  export let categories: Category[];
  export let addIncome: any;
  export let getIncomes: any;
</script>

<form
  id="form"
  hidden
  on:submit|preventDefault={async () => {
    await addIncome(info);
    await getIncomes();
  }}
>
  <div>
    <label for="title">Enter a title</label><br />
    <input
      type="text"
      id="title"
      class="border-4"
      bind:value={info.title}
      placeholder="My Title"
    />
  </div>
  <br />
  <div>
    <label for="description">Enter a description(optional)</label><br />
    <input
      type="text"
      id="description"
      class="border-4"
      bind:value={info.description}
      placeholder="My Description"
    />
  </div>
  <br />
  <div>
    <label for="currency">Enter a currency</label><br />
    <select id="currency" bind:value={info.currency}>
      {#each CURRENCIES as currency}
        <option value={currency}>{currency}</option>
      {/each}
    </select>
  </div>
  <br />
  <div>
    <label for="amount">Enter an amount</label>
    <input
      class="border-4"
      type="text"
      id="amount"
      bind:value={info.amount}
      on:input={checkIfValidNumberInput}
    />
  </div>
  <br />
  <div>
    <label for="payment">Enter a payment method</label><br />
    <select id="payment" bind:value={info.paymentmethod}>
      {#each paymentMethods as method}
        <option value={method}>{method}</option>
      {/each}
    </select>
  </div>
  <br />
  <div>
    <label for="client">Enter an client(who's paying)</label>
    <input
      class="border-4"
      type="text"
      id="client"
      bind:value={info.participant}
    />
  </div>
  <br />
  <div>
    <label for="recurring">Is recurring</label>
    <input
      type="checkbox"
      id="recurring"
      on:change={() => {
        info.recurring = !info.recurring;
      }}
    />
    <!--TODO: make a data input-->
    {#if info.recurring}
      <label for="interval">Enter The Interval</label>
      <input
        type="text"
        id="interval"
        class="border-4"
        placeholder="Interval"
        bind:value={info.interval}
        on:input={checkIfValidNumberInput}
      />
    {/if}
  </div>
  <br />
  <div>
    <label for="category">Enter a category</label><br />
    <select id="category" bind:value={info.category}>
      {#each categories as category}
        {#if category.type == "income"}
          <option value={category.title}>{category.title}</option>
        {/if}
      {/each}
    </select>
  </div>
  <br />
  <div>
    <label for="isTaxxed">Is taxxed</label>
    <input
      type="checkbox"
      id="isTaxxed"
      on:change={() => {
        info.taxxed = !info.taxxed;
      }}
    />
    {#if info.taxxed}
      <label for="tax">Enter a tax</label>
      <input
        class="border-4"
        type="text"
        id="tax"
        bind:value={info.tax}
        on:input={checkIfValidNumberInput}
      />
    {/if}
  </div>
  <br />
  <div>
    <label for="received">Received</label>
    <input
      checked
      type="checkbox"
      id="received"
      on:change={() => {
        info.fulfilled = !info.fulfilled;
      }}
    />
  </div>
  <br />
  <button type="submit">Submit</button>
</form>

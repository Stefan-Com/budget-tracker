<script lang="ts">
  import { checkIfValidNumberInput, CURRENCIES } from "../../utils";
  export let addCategory: any;
  export let getCategories: any;
  let info: Category = {
    id: null,
    parentid: null,
    title: "",
    description: "",
    currency: CURRENCIES[0],
    budget: 0,
    spent: 0,
    gotten: 0,
    type: "expense",
    budgeted: false,
  };
</script>

<form
  id="form"
  hidden
  on:submit|preventDefault={async () => {
    await addCategory(info);
    await getCategories();
  }}
>
  <div>
    <label for="title">Enter a title</label><br />
    <input
      type="text"
      id="title"
      class="border-4"
      bind:value={info.title}
      placeholder="My Category"
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
    <div>
      <label for="type">Enter a type(income/expense)</label><br />
      <select id="type" bind:value={info.type}>
        <option value="income">Income</option>
        <option value="expense" selected>Expense</option>
      </select>
    </div>
    <br />
    {#if info.type == "expense"}
      <div>
        <label for="budgeted">Is budgeted</label>
        <input
          type="checkbox"
          id="budgeted"
          on:change={() => {
            info.budgeted = !info.budgeted;
          }}
        />
        {#if info.budgeted == true}
          <label for="budget">Enter The Budget</label>
          <input
            type="text"
            id="budget"
            class="border-4"
            placeholder="Budget"
            bind:value={info.budget}
            on:input={checkIfValidNumberInput}
          />
        {/if}
      </div>
      <div>
        <label for="spent">Enter The Already Spent Amount</label>
        <input
          type="text"
          id="spent"
          class="border-4"
          placeholder="Spent"
          bind:value={info.spent}
          on:input={checkIfValidNumberInput}
        />
      </div>
    {:else}
      <div>
        <label for="gotten">Enter The Already Gotten Amount</label>
        <input
          type="text"
          id="gotten"
          class="border-4"
          placeholder="Gotten"
          bind:value={info.gotten}
          on:input={checkIfValidNumberInput}
        />
      </div>{/if}
  </div>
  <br />
  <button type="submit">Submit</button>
</form>

<script lang="ts">
  import { onMount } from "svelte";
  import {
    checkIfValidNumberInput,
    CURRENCIES,
    getCategories,
  } from "../../utils";
  import CategoryForm from "../../components/forms/CategoryForm.svelte";
  import EditBtn from "../../components/buttons/EditBtn.svelte";
  import DeleteBtn from "../../components/buttons/DeleteBtn.svelte";
  import InfoBtn from "../../components/buttons/InfoBtn.svelte";
  import { addCategory, deleteCategory, editCategory } from "../../controllers";
  let showCategories = false;
  let form: any;

  let currentCategory: Category | null;
  let currentEditCategory: Category | null;
  let categories: any;

  onMount(async () => {
    form = document.getElementById("form");
    categories = await getCategories();
    if (!categories) categories = [];
    showCategories = true;
  });

  function displayInfo(info: Category) {
    currentCategory = info;
  }
  function editInfo(info: Category) {
    currentEditCategory = info;
  }
</script>

<div>
  <div>
    <h1 class="pl-5">Your Categories:</h1>
    <ul>
      {#if showCategories}
        {#each categories as category}
          <li
            class="hover:bg-stone-400 flex justify-between w-1/2 p-5 rounded-3xl bg-stone-200"
          >
            <div>
              Name: {category.title}
              <br />
              Type: {category.type}
            </div>
            <div class="flex justify-around items-center w-28">
              <EditBtn
                handleClick={() => {
                  editInfo(category);
                }}
              />
              <DeleteBtn
                handleClick={async () => {
                  await deleteCategory(category);
                  categories = await getCategories();
                  if (!categories) categories = [];
                }}
              />
              <InfoBtn
                handleClick={() => {
                  displayInfo(category);
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
          }}>Add New Category</button
        >
      </li>
    </ul>
    {#if currentCategory}
      <dialog
        data-te-toggle="modal"
        id="modal"
        open
        class="border-4 border-black top-0 left-0 w-96 h-fit"
      >
        <div>
          <h3>Title: {currentCategory.title}</h3>
        </div>
        <br />
        <div>
          <h3>Description: {currentCategory.description}</h3>
        </div>
        <br />
        <div>
          <h3>Currency: {currentCategory.currency}</h3>
        </div>
        <br />
        <div>
          <h3>Type: {currentCategory.type}</h3>
          {#if currentCategory.type == "expense"}
            <h3>Budgeted: {!!currentCategory.budgeted ? "Yes" : "No"}</h3>
            {#if currentCategory.budgeted}
              <div>
                <h3>Budget: {currentCategory.budget}</h3>
              </div>
            {/if}
            <div>
              <h3>Spent: {currentCategory.spent}</h3>
            </div>
          {:else}
            <h3>Gotten: {currentCategory.gotten}</h3>
          {/if}
        </div>
        <br />
        <div>
          <button
            on:click={() => {
              currentCategory = null;
            }}>Exit</button
          >
        </div>
        <br />
      </dialog>
    {:else if currentEditCategory}
      <dialog
        data-te-toggle="modal"
        id="modal"
        open
        class="border-4 border-black top-0 left-0 w-96 h-fit"
      >
        <div>
          <h3>
            Title: <input bind:value={currentEditCategory.title} />
          </h3>
        </div>
        <br />
        <div>
          <h3>
            Description: <input bind:value={currentEditCategory.description} />
          </h3>
        </div>
        <br />
        <div>
          <h3>
            Currency: <select bind:value={currentEditCategory.currency}>
              {#each CURRENCIES as currency}
                <option value={currency}>{currency}</option>
              {/each}
            </select>
          </h3>
        </div>
        <br />
        <div>
          <h3>
            Type: <select bind:value={currentEditCategory.type}>
              <option selected value={currentEditCategory.type}
                >{currentEditCategory.type}</option
              >
              <option
                value={currentEditCategory.type == "income"
                  ? "expense"
                  : "income"}
                >{currentEditCategory.type == "income"
                  ? "expense"
                  : "income"}</option
              >
            </select>
          </h3>
          {#if currentEditCategory.type == "expense"}
            <div>
              <h3>
                Budgeted: <input
                  type="checkbox"
                  bind:checked={currentEditCategory.budgeted}
                />
              </h3>
              {#if currentEditCategory.budgeted}
                <h3>
                  Budget: <input
                    bind:value={currentEditCategory.budget}
                    on:input={checkIfValidNumberInput}
                  />
                </h3>
              {/if}
              <h3>
                Spent: <input
                  bind:value={currentEditCategory.spent}
                  on:input={checkIfValidNumberInput}
                />
              </h3>
            </div>
          {:else}
            <h3>
              Amount Gotten: <input
                bind:value={currentEditCategory.gotten}
                on:input={checkIfValidNumberInput}
              />
            </h3>
          {/if}
        </div>
        <br />
        <div>
          <button
            on:click={async () => {
              await editCategory(currentEditCategory);
              currentEditCategory = null;
              categories = await getCategories();
              if (!categories) categories = [];
            }}>Save changes</button
          >
          <button
            on:click={async () => {
              currentEditCategory = null;
            }}>Exit</button
          >
        </div>
        <br />
      </dialog>
    {/if}
  </div>
  <CategoryForm
    {addCategory}
    getCategories={() => {
      return new Promise(async (resolve, reject) => {
        categories = await getCategories();
        if (!categories) categories = [];
        resolve(null);
      });
    }}
  />
</div>

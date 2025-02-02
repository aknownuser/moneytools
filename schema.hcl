schema "public" {}

table "accounts" {
  schema = schema.public
  
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "balance" {
    null = false
    type = decimal(10,2)
    default = 0
  }
  primary_key {
    columns = [column.id]
  }
  index "accounts_deleted_at_idx" {
    columns = [column.deleted_at]
  }
}

table "transactions" {
  schema = schema.public
  
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "amount" {
    null = false
    type = decimal(10,2)
  }
  column "date" {
    null = false
    type = timestamptz
  }
  column "category_id" {
    null = false
    type = bigint
  }
  column "account_id" {
    null = false
    type = bigint
  }
  
  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_transactions_category" {
    columns = [column.category_id]
    ref_columns = [table.categories.column.id]
    on_delete = "CASCADE"
  }

  foreign_key "fk_transactions_account" {
    columns = [column.account_id]
    ref_columns = [table.accounts.column.id]
    on_delete = "CASCADE"
  }

  index "transactions_deleted_at_idx" {
    columns = [column.deleted_at]
  }

  
}

table "categories" {
  schema = schema.public
  
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.id]
  }
  index "categories_deleted_at_idx" {
    columns = [column.deleted_at]
  }
}

table "budgets" {
  schema = schema.public
  
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "category_id" {
    null = false
    type = bigint
  }
  column "amount" {
    null = false
    type = decimal(10,2)
  }
  column "account_id" {
    null = false
    type = bigint
  }
  column "date" {
    null = false
    type = timestamptz
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_budgets_category" {
    columns = [column.category_id]
    ref_columns = [table.categories.column.id]
    on_delete = "CASCADE"
   }

  foreign_key "fk_budgets_account" {
    columns = [column.account_id]
    ref_columns = [table.accounts.column.id]
    on_delete = "CASCADE"
  }

  index "budgets_deleted_at_idx" {
    columns = [column.deleted_at]
  }
}

table "budget_changes" {
  schema = schema.public
  
  column "id" {
    null = false
    type = bigserial
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "budget_id" {
    null = false
    type = bigint
  }
  column "amount" {
    null = false
    type = decimal(10,2)
  }
  column "reason" {
    null = false
    type = text
  }
  column "date" {
    null = false
    type = timestamptz
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_budget_changes_budget" {
    columns = [column.budget_id]
    ref_columns = [table.budgets.column.id]
    on_delete = "CASCADE"
  }

  index "budget_changes_deleted_at_idx" {
    columns = [column.deleted_at]
  }
}
import { Pencil, ArrowUpDown, Trash2 } from "lucide-react";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { models } from "../../wailsjs/go/models";
import { useState, useEffect } from "react";

type Product = models.Product;

interface ProductTableProps {
  products: Product[];
  selectedProducts: Record<number, boolean>;
  onSelect: (id: number, isSelected: boolean) => void;
  onEdit: (product: Product) => void;
  onDelete: () => void;
}

export function ProductTable({
  products,
  selectedProducts,
  onSelect,
  onEdit,
  onDelete,
}: ProductTableProps) {
  const [sortDirection, setSortDirection] = useState<'asc' | 'desc'>(() => {
    const savedSortDirection = localStorage.getItem('sortDirection');
    return savedSortDirection ? (savedSortDirection as 'asc' | 'desc') : 'asc';
  });
  
  // Сохранение состояния сортировки при изменении
  useEffect(() => {
    localStorage.setItem('sortDirection', sortDirection);
  }, [sortDirection]);
  
  const toggleSort = () => {
    setSortDirection(prev => prev === 'asc' ? 'desc' : 'asc');
  };
  
  const sortedProducts = [...products].sort((a, b) => {
    if (sortDirection === 'asc') {
      return a.id - b.id;
    } else {
      return b.id - a.id;
    }
  });
  
  const handleRowClick = (product: Product) => {
    onSelect(product.id, !selectedProducts[product.id]);
  };

  // Проверяем, есть ли выбранные продукты
  const hasSelectedProducts = Object.values(selectedProducts).some(value => value);

  return (
    <div className="border rounded-md flex flex-col h-full">
      <div className="overflow-auto flex-1">
        <table className="w-full border-collapse">
          <thead className="sticky-header">
            <tr>
              <th className="w-[50px] px-4 py-2 text-center font-medium text-muted-foreground">
                <Button
                  variant={hasSelectedProducts ? "destructive" : "ghost"}
                  size="icon"
                  onClick={onDelete}
                  disabled={!hasSelectedProducts}
                  className={`h-7 w-7 ${!hasSelectedProducts ? 'text-gray-300 hover:text-gray-300 hover:bg-transparent' : ''}`}
                >
                  <Trash2 className="h-4 w-4" />
                </Button>
              </th>
              <th className="px-4 py-2 text-left font-medium text-muted-foreground">
                <div className="flex items-center">
                  Наименование
                  <Button variant="ghost" size="sm" onClick={toggleSort} className="ml-2 h-7 w-7 p-0">
                    <ArrowUpDown className="h-4 w-4" />
                  </Button>
                </div>
              </th>
              <th className="w-[150px] px-4 py-2 text-center font-medium text-muted-foreground whitespace-nowrap">Время обработки</th>
              <th className="w-[80px] px-4 py-2 text-center font-medium text-muted-foreground">Действия</th>
            </tr>
          </thead>
          <tbody>
            {!products || products.length === 0 ? (
              <tr>
                <td colSpan={4} className="text-center py-8 text-lg">
                  Нет данных для отображения
                </td>
              </tr>
            ) : (
              sortedProducts.map((product) => (
                <tr 
                  key={product.id} 
                  className="border-b cursor-pointer hover:bg-gray-50"
                  onClick={() => handleRowClick(product)}
                >
                  <td className="px-4 py-3" onClick={(e) => e.stopPropagation()}>
                    <Checkbox
                      checked={selectedProducts[product.id] || false}
                      onCheckedChange={(checked) =>
                        onSelect(product.id, checked === true)
                      }
                    />
                  </td>
                  <td className="px-4 py-3 font-medium">{product.name}</td>
                  <td className="px-4 py-3 text-center">{product.processingTime.toFixed(2)} ч.</td>
                  <td className="px-4 py-3">
                    <div className="flex justify-center">
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={(e) => {
                          e.stopPropagation();
                          onEdit(product);
                        }}
                        className="h-7 w-7"
                      >
                        <Pencil className="h-4 w-4" />
                      </Button>
                    </div>
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
} 
import { useEffect, useState } from "react";

// Декларация для Wails runtime
declare global {
  interface Window {
    runtime?: unknown;
    go?: unknown;
  }
}
import { GetProducts, SearchProducts, DeleteProducts } from "../wailsjs/go/main/App";
import { ProductTable } from "./components/ProductTable";
import { AddProductDialog } from "./components/AddProductDialog";
import { EditProductDialog } from "./components/EditProductDialog";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { ToastProvider } from "./components/ui/toast";
import { Toaster } from "./components/Toaster";
import { useToast } from "./hooks/use-toast";
import { models } from "../wailsjs/go/models";
import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle } from "./components/ui/alert-dialog";
import { Switch } from "./components/ui/switch";
import { Label } from "./components/ui/label";

type Product = models.Product;

// Функция ожидания готовности Wails runtime с таймаутом
const waitForWailsRuntime = (): Promise<void> => {
  return new Promise((resolve) => {
    // Wails использует window.go для биндингов Go функций
    if (window.go) {
      resolve();
      return;
    }
    
    let attempts = 0;
    const maxAttempts = 100; // 5 секунд максимум (100 * 50ms)
    
    const checkRuntime = () => {
      attempts++;
      if (window.go || attempts >= maxAttempts) {
        resolve();
      } else {
        setTimeout(checkRuntime, 50);
      }
    };
    checkRuntime();
  });
};

function App() {
  const [products, setProducts] = useState<Product[]>([]);
  const [isReady, setIsReady] = useState(false);
  const [searchQuery, setSearchQuery] = useState(() => {
    const savedSearchQuery = localStorage.getItem('searchQuery');
    return savedSearchQuery || "";
  });
  const [selectedProducts, setSelectedProducts] = useState<Record<number, boolean>>({});
  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false);
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false);
  const [currentProduct, setCurrentProduct] = useState<Product | null>(null);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [selectedIdsToDelete, setSelectedIdsToDelete] = useState<number[]>([]);
  const [filterBySelected, setFilterBySelected] = useState(() => {
    const savedFilter = localStorage.getItem('filterBySelected');
    return savedFilter ? JSON.parse(savedFilter) : false;
  });
  const { toast } = useToast();

  // Ожидание готовности Wails runtime и загрузка продуктов
  useEffect(() => {
    const init = async () => {
      await waitForWailsRuntime();
      setIsReady(true);
      
      // Если есть сохраненный поисковый запрос, выполняем поиск
      if (searchQuery) {
        handleSearch(searchQuery);
      } else {
        loadProducts();
      }
    };
    init();
  }, []);

  // Сохранение состояния фильтрации при изменении
  useEffect(() => {
    localStorage.setItem('filterBySelected', JSON.stringify(filterBySelected));
  }, [filterBySelected]);

  // Сохранение поискового запроса при изменении
  useEffect(() => {
    localStorage.setItem('searchQuery', searchQuery);
  }, [searchQuery]);

  // Загрузка всех продуктов
  const loadProducts = async () => {
    try {
      const data = await GetProducts();
      // Убедимся, что data - это массив
      const productsArray = Array.isArray(data) ? data : [];
      setProducts(productsArray);
      
      // Загружаем сохраненные выбранные продукты
      const savedSelectedProducts = localStorage.getItem('selectedProducts');
      if (savedSelectedProducts && productsArray.length > 0) {
        const parsedSelectedProducts = JSON.parse(savedSelectedProducts);
        // Фильтруем сохраненные ID, чтобы убедиться, что они существуют в текущих данных
        const validSelectedProducts: Record<number, boolean> = {};
        productsArray.forEach(product => {
          if (parsedSelectedProducts[product.id]) {
            validSelectedProducts[product.id] = true;
          }
        });
        setSelectedProducts(validSelectedProducts);
      } else {
        setSelectedProducts({});
      }
    } catch (error) {
      console.error("Ошибка загрузки продуктов:", error);
      setProducts([]);
      toast({
        title: "Ошибка",
        description: "Не удалось загрузить данные",
        variant: "destructive",
      });
    }
  };

  // Поиск продуктов
  const handleSearch = async (query: string) => {
    try {
      setSearchQuery(query);
      // Проверка на специальные символы или потенциально проблемные запросы
      const safeQuery = query.trim();
      const data = await SearchProducts(safeQuery);
      
      // Убедимся, что data - это массив (даже если пустой)
      if (Array.isArray(data)) {
        setProducts(data);
      } else {
        // Если данные не являются массивом, используем пустой массив
        setProducts([]);
        console.error("Получены некорректные данные:", data);
      }
    } catch (error) {
      console.error("Ошибка при поиске:", error);
      // Очистка поиска при ошибке
      setSearchQuery("");
      toast({
        title: "Ошибка",
        description: "Ошибка при поиске. Попробуйте другой запрос.",
        variant: "destructive",
      });
      // При ошибке загружаем все продукты
      loadProducts();
    }
  };

  // Подготовка к удалению выбранных продуктов
  const prepareDeleteSelected = () => {
    const selectedIds = Object.entries(selectedProducts)
      .filter(([_, isSelected]) => isSelected)
      .map(([id]) => parseInt(id));

    if (selectedIds.length === 0) {
      toast({
        title: "Внимание",
        description: "Не выбрано ни одной записи",
      });
      return;
    }

    setSelectedIdsToDelete(selectedIds);
    setIsDeleteDialogOpen(true);
  };

  // Удаление выбранных продуктов
  const handleDeleteSelected = async () => {
    try {
      await DeleteProducts(selectedIdsToDelete);
      toast({
        title: "Успешно",
        description: `Удалено записей: ${selectedIdsToDelete.length}`,
      });
      setIsDeleteDialogOpen(false);
      loadProducts();
    } catch (error) {
      toast({
        title: "Ошибка",
        description: "Не удалось удалить записи",
        variant: "destructive",
      });
    }
  };

  // Обработка выбора продукта
  const handleSelectProduct = (id: number, isSelected: boolean) => {
    const newSelectedProducts = {
      ...selectedProducts,
      [id]: isSelected,
    };
    setSelectedProducts(newSelectedProducts);
    
    // Сохраняем выбранные продукты в localStorage
    localStorage.setItem('selectedProducts', JSON.stringify(newSelectedProducts));
  };

  // Снятие выделения со всех продуктов
  const handleClearSelection = () => {
    setSelectedProducts({});
    localStorage.removeItem('selectedProducts');
  };

  // Переключение фильтра по выбранным записям
  const toggleFilterBySelected = () => {
    setFilterBySelected(!filterBySelected);
  };

  // Открытие диалога редактирования
  const handleEdit = (product: Product) => {
    setCurrentProduct(product);
    setIsEditDialogOpen(true);
  };

  // Обработка успешного добавления/редактирования
  const handleSuccess = () => {
    loadProducts();
  };

  // Фильтрация продуктов
  const filteredProducts = filterBySelected && Object.values(selectedProducts).some(value => value)
    ? products.filter(product => selectedProducts[product.id])
    : products;

  // Показываем загрузку пока Wails runtime не готов
  if (!isReady) {
    return (
      <div className="flex items-center justify-center h-screen">
        <div className="text-lg">Загрузка...</div>
      </div>
    );
  }

  return (
    <ToastProvider>
      <div className="flex flex-col h-screen overflow-hidden">
        <div className="sticky top-0 bg-white z-20 border-b shadow-sm">
          <div className="container mx-auto py-3 px-4">
            <h1 className="text-2xl font-bold mb-3">Редактор базы данных</h1>
            
            <div className="flex flex-col md:flex-row gap-3 mb-1">
              <Input
                placeholder="Поиск по наименованию..."
                value={searchQuery}
                onChange={(e) => handleSearch(e.target.value)}
                className="md:w-1/3"
              />
              <div className="flex items-center gap-2 ml-2">
                <Switch 
                  id="filter-selected" 
                  checked={filterBySelected}
                  onCheckedChange={toggleFilterBySelected}
                />
                <Label htmlFor="filter-selected">Показать выбранные</Label>
              </div>
              <div className="flex gap-2 ml-auto">
                <Button onClick={() => setIsAddDialogOpen(true)}>
                  Добавить запись
                </Button>
                <Button variant="outline" onClick={handleClearSelection}>
                  Снять выделение
                </Button>
              </div>
            </div>
          </div>
        </div>

        <div className="flex-1 overflow-hidden pb-4">
          <div className="container mx-auto px-4 pt-4 h-full">
            <ProductTable
              products={filteredProducts}
              selectedProducts={selectedProducts}
              onSelect={handleSelectProduct}
              onEdit={handleEdit}
              onDelete={prepareDeleteSelected}
            />
          </div>
        </div>

        <AddProductDialog
          isOpen={isAddDialogOpen}
          onClose={() => setIsAddDialogOpen(false)}
          onSuccess={handleSuccess}
        />

        {currentProduct && (
          <EditProductDialog
            isOpen={isEditDialogOpen}
            onClose={() => setIsEditDialogOpen(false)}
            product={currentProduct}
            onSuccess={handleSuccess}
          />
        )}

        <AlertDialog open={isDeleteDialogOpen} onOpenChange={setIsDeleteDialogOpen}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Подтверждение удаления</AlertDialogTitle>
              <AlertDialogDescription>
                Вы уверены, что хотите удалить выбранные записи ({selectedIdsToDelete.length})? Это действие нельзя отменить.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setIsDeleteDialogOpen(false)}>Отмена</AlertDialogCancel>
              <AlertDialogAction onClick={handleDeleteSelected}>Удалить</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>

        <Toaster />
      </div>
    </ToastProvider>
  );
}

export default App; 
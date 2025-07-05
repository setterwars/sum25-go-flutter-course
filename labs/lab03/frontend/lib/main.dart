import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'screens/chat_screen.dart';
import 'services/api_service.dart';
import 'models/message.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Wrap MaterialApp with MultiProvider to provide ApiService and ChatProvider
    return MultiProvider(
      providers: [
        Provider<ApiService>(
          create: (_) => ApiService(),
          dispose: (_, apiService) => apiService.dispose(),
        ),
        ChangeNotifierProvider<ChatProvider>(
          create: (context) => ChatProvider(context.read<ApiService>()),
        ),
      ],
      child: MaterialApp(
        title: 'Lab 03 REST API Chat',
        theme: ThemeData(
          // Customize theme colors
          primarySwatch: Colors.blue,
          colorScheme: ColorScheme.fromSeed(
            seedColor: Colors.blue,
            secondary: Colors.orange, // HTTP cat theme accent color
          ),
          // Configure app bar theme
          appBarTheme: const AppBarTheme(
            backgroundColor: Colors.blue,
            foregroundColor: Colors.white,
            elevation: 2,
          ),
          // Configure elevated button theme
          elevatedButtonTheme: ElevatedButtonThemeData(
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.blue,
              foregroundColor: Colors.white,
            ),
          ),
          useMaterial3: true,
        ),
        home: const ChatScreen(),
        // Add error handling for navigation
        onGenerateRoute: (settings) {
          // Handle unknown routes
          return MaterialPageRoute(
            builder: (context) => const ChatScreen(),
          );
        },
      ),
    );
  }
}

// Provider class for managing app state
class ChatProvider extends ChangeNotifier {
  final ApiService _apiService;
  List<Message> _messages = [];
  bool _isLoading = false;
  String? _error;

  // Constructor that takes ApiService
  ChatProvider(this._apiService);

  // Getters for all private fields
  List<Message> get messages => _messages;
  bool get isLoading => _isLoading;
  String? get error => _error;

  // Load messages from API
  Future<void> loadMessages() async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      _messages = await _apiService.getMessages();
      _error = null;
    } catch (e) {
      _error = e.toString();
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  // Create a new message
  Future<void> createMessage(CreateMessageRequest request) async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      final newMessage = await _apiService.createMessage(request);
      _messages.add(newMessage);
      _error = null;
    } catch (e) {
      _error = e.toString();
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  // Update an existing message
  Future<void> updateMessage(int id, UpdateMessageRequest request) async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      final updatedMessage = await _apiService.updateMessage(id, request);
      final index = _messages.indexWhere((message) => message.id == id);
      if (index != -1) {
        _messages[index] = updatedMessage;
      }
      _error = null;
    } catch (e) {
      _error = e.toString();
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  // Delete a message
  Future<void> deleteMessage(int id) async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      await _apiService.deleteMessage(id);
      _messages.removeWhere((message) => message.id == id);
      _error = null;
    } catch (e) {
      _error = e.toString();
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  // Refresh messages by clearing and reloading
  Future<void> refreshMessages() async {
    _messages.clear();
    await loadMessages();
  }

  // Clear error state
  void clearError() {
    _error = null;
    notifyListeners();
  }
}
